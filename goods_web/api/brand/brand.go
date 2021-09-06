package brand

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/goods_web/api"
	"github.com/xlt/shop_web/goods_web/forms"
	"github.com/xlt/shop_web/goods_web/global"
	"github.com/xlt/shop_web/goods_web/proto"
)

func BrandList(ctx *gin.Context) {
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	rsp, err := global.GoodsClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       int32(pnInt),
		PagePerNums: int32(pSizeInt),
	})

	if err != nil {
		zap.S().Errorw("global.GoodsClient.BrandList failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	result := make([]interface{}, 0)
	reMap := make(map[string]interface{})
	reMap["total"] = rsp.Total
	for _, value := range rsp.Data[pnInt : pnInt*pSizeInt+pSizeInt] {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["name"] = value.Name
		reMap["logo"] = value.Logo

		result = append(result, reMap)
	}
	reMap["data"] = result

	ctx.JSON(http.StatusOK, reMap)
}

func NewBrand(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(&brandForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	rsp, err := global.GoodsClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.CreateBrand failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	request := make(map[string]interface{})
	request["id"] = rsp.Id
	request["name"] = rsp.Name
	request["logo"] = rsp.Logo

	ctx.JSON(http.StatusOK, request)
}

func DeleteBrand(ctx *gin.Context) {
	brandID := ctx.Param("id")
	brandIDInt, err := strconv.Atoi(brandID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}
	_, err = global.GoodsClient.DeleteBrand(context.Background(), &proto.BrandRequest{Id: int32(brandIDInt)})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.DeleteBrand failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func UpdateBrand(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(&brandForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	brandID := ctx.Param("id")
	brandIDInt, err := strconv.Atoi(brandID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	_, err = global.GoodsClient.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   int32(brandIDInt),
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.UpdateBrand failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func GetCategoryBrandList(ctx *gin.Context) {
	brandID := ctx.Param("id")
	brandIDInt, err := strconv.Atoi(brandID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	rsp, err := global.GoodsClient.GetCategoryBrandList(context.Background(), &proto.CategoryInfoRequest{
		Id: int32(brandIDInt),
	})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.GetCategoryBrandList( failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["name"] = value.Name
		reMap["logo"] = value.Logo

		result = append(result, reMap)
	}

	ctx.JSON(http.StatusOK, result)
}

func CategoryBrandList(ctx *gin.Context) {
	//所有的list返回的数据结构
	/*
		{
			"total": 100,
			"data":[{},{}]
		}
	*/
	rsp, err := global.GoodsClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.CategoryBrandList failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	reMap := map[string]interface{}{
		"total": rsp.Total,
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["category"] = map[string]interface{}{
			"id":   value.Category.Id,
			"name": value.Category.Name,
		}
		reMap["brand"] = map[string]interface{}{
			"id":   value.Brand.Id,
			"name": value.Brand.Name,
			"logo": value.Brand.Logo,
		}

		result = append(result, reMap)
	}

	reMap["data"] = result
	ctx.JSON(http.StatusOK, reMap)
}

func NewCategoryBrand(ctx *gin.Context) {
	categoryBrandForm := forms.CategoryBrandForm{}
	if err := ctx.ShouldBindJSON(&categoryBrandForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	rsp, err := global.GoodsClient.CreateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.CreateCategoryBrand failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	response := make(map[string]interface{})
	response["id"] = rsp.Id

	ctx.JSON(http.StatusOK, response)
}

func UpdateCategoryBrand(ctx *gin.Context) {
	categoryBrandForm := forms.CategoryBrandForm{}
	if err := ctx.ShouldBindJSON(&categoryBrandForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	brandID := ctx.Param("id")
	brandIDInt, err := strconv.Atoi(brandID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	_, err = global.GoodsClient.UpdateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id:         int32(brandIDInt),
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.UpdateCategoryBrand failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func DeleteCategoryBrand(ctx *gin.Context) {
	brandID := ctx.Param("id")
	brandIDInt, err := strconv.Atoi(brandID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}
	_, err = global.GoodsClient.DeleteCategoryBrand(context.Background(), &proto.CategoryBrandRequest{Id: int32(brandIDInt)})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.DeleteCategoryBrand failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, "")
}
