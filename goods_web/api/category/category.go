package category

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/xlt/shop_web/goods_web/api"
	"github.com/xlt/shop_web/goods_web/forms"
	"github.com/xlt/shop_web/goods_web/global"
	"github.com/xlt/shop_web/goods_web/proto"
)

func List(ctx *gin.Context) {
	r, err := global.GoodsClient.GetAllCategorysList(context.Background(), &empty.Empty{})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.GetAllCategorysList failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(r.JsonData), &data)
	if err != nil {
		zap.S().Errorw("json.Unmarshal failed ", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func Detail(ctx *gin.Context) {
	goodsID := ctx.Param("id")
	goodsIDInt, err := strconv.Atoi(goodsID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	reMap := make(map[string]interface{})
	subCategorys := make([]interface{}, 0)
	if r, err := global.GoodsClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: int32(goodsIDInt),
	}); err != nil {
		zap.S().Errorw("global.GoodsClient.GetSubCategory failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	} else {
		//写文档 特别是数据多的时候很慢， 先开发后写文档
		for _, value := range r.SubCategorys {
			subCategorys = append(subCategorys, map[string]interface{}{
				"id":              value.Id,
				"name":            value.Name,
				"level":           value.Level,
				"parent_category": value.ParentCategory,
				"is_tab":          value.IsTab,
			})
		}
		reMap["id"] = r.Info.Id
		reMap["name"] = r.Info.Name
		reMap["level"] = r.Info.Level
		reMap["parent_category"] = r.Info.ParentCategory
		reMap["is_tab"] = r.Info.IsTab
		reMap["sub_categorys"] = subCategorys

		ctx.JSON(http.StatusOK, reMap)
	}
	return
}

func New(ctx *gin.Context) {
	categoryForm := forms.CategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	rsp, err := global.GoodsClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:           categoryForm.Name,
		IsTab:          *categoryForm.IsTab,
		Level:          categoryForm.Level,
		ParentCategory: categoryForm.ParentCategory,
	})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.CreateCategory failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	request := make(map[string]interface{})
	request["id"] = rsp.Id
	request["name"] = rsp.Name
	request["parent"] = rsp.ParentCategory
	request["level"] = rsp.Level
	request["is_tab"] = rsp.IsTab

	ctx.JSON(http.StatusOK, request)
}

func Delete(ctx *gin.Context) {
	goodsID := ctx.Param("id")
	goodsIDInt, err := strconv.Atoi(goodsID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	//1. 先查询出该分类写的所有子分类
	//2. 将所有的分类全部逻辑删除
	//3. 将该分类下的所有的商品逻辑删除
	_, err = global.GoodsClient.DeleteCategory(context.Background(), &proto.DeleteCategoryRequest{Id: int32(goodsIDInt)})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.DeleteCategory failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func Update(ctx *gin.Context) {
	categoryForm := forms.UpdateCategoryForm{}
	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	goodsID := ctx.Param("id")
	goodsIDInt, err := strconv.Atoi(goodsID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	request := &proto.CategoryInfoRequest{
		Id:   int32(goodsIDInt),
		Name: categoryForm.Name,
	}
	if categoryForm.IsTab != nil {
		request.IsTab = *categoryForm.IsTab
	}
	_, err = global.GoodsClient.UpdateCategory(context.Background(), request)
	if err != nil {
		zap.S().Errorw("global.GoodsClient.UpdateCategory failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
