package goods

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/xlt/shop_web/goods_web/forms"
	"github.com/xlt/shop_web/goods_web/global"
	"github.com/xlt/shop_web/goods_web/proto"
)

func List(ctx *gin.Context) {
	request := &proto.GoodsFilterRequest{}
	var rowCount int

	priceMin := ctx.DefaultQuery("pmin", "0")
	priceMinInt, _ := strconv.Atoi(priceMin)
	request.PriceMin = int32(priceMinInt)

	priceMax := ctx.DefaultQuery("pmax", "0")
	priceMaxInt, _ := strconv.Atoi(priceMax)
	request.PriceMin = int32(priceMaxInt)

	isHot := ctx.DefaultQuery("ih", "0")
	if isHot == "1" {
		request.IsHot = true
	}
	isNew := ctx.DefaultQuery("in", "0")
	if isNew == "1" {
		request.IsNew = true
	}
	isTab := ctx.DefaultQuery("it", "0")
	if isTab == "1" {
		request.IsTab = true
	}

	topCategoryId := ctx.DefaultQuery("c", "0")
	topCategoryIdInt, _ := strconv.Atoi(topCategoryId)
	request.TopCategory = int32(topCategoryIdInt)

	page := ctx.DefaultQuery("p", "0")
	pageInt, _ := strconv.Atoi(page)
	request.Pages = int32(pageInt)

	pageNum := ctx.DefaultQuery("pnum", "0")
	pageNumInt, _ := strconv.Atoi(pageNum)
	request.PagePerNums = int32(pageNumInt)

	keywords := ctx.DefaultQuery("q", "")
	request.KeyWords = keywords

	brandId := ctx.DefaultQuery("b", "0")
	brandIdInt, _ := strconv.Atoi(brandId)
	request.Brand = int32(brandIdInt)

	response, err := global.GoodsClient.GoodsList(context.Background(), request)
	if err != nil {
		zap.S().Errorw("global.GoodsClient.GoodsList failed", "msg", err.Error())
		HandleGrpcErrorToHttp(ctx, err)
	}
	rowCount = len(response.Data)

	goodsList := make([]interface{}, response.Total)
	for i := 0; i < rowCount; i++ {
		goodsList[i] = map[string]interface{}{
			"id":          response.Data[i].Id,
			"name":        response.Data[i].Name,
			"goods_brief": response.Data[i].GoodsBrief,
			"desc":        response.Data[i].GoodsDesc,
			"ship_free":   response.Data[i].ShipFree,
			"images":      response.Data[i].Images,
			"desc_images": response.Data[i].DescImages,
			"front_image": response.Data[i].GoodsFrontImage,
			"shop_price":  response.Data[i].ShopPrice,
			"category": map[string]interface{}{
				"id":   response.Data[i].Category.Id,
				"name": response.Data[i].Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   response.Data[i].Brand.Id,
				"name": response.Data[i].Brand.Name,
				"logo": response.Data[i].Brand.Logo,
			},
			"is_hot":  response.Data[i].IsHot,
			"is_new":  response.Data[i].IsNew,
			"on_sale": response.Data[i].OnSale,
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total": response.Total,
		"data":  goodsList,
	})
}

func New(ctx *gin.Context) {
	goodsForm := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&goodsForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	rsp, err := global.GoodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.Brand,
	})
	if err != nil {
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

func Detail(ctx *gin.Context) {
	goodsID := ctx.Param("id")
	goodsIDInt, err := strconv.Atoi(goodsID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	response, err := global.GoodsClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{Id: int32(goodsIDInt)})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.GetGoodsDetail failed", "msg", err.Error())
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	goodsInfo := map[string]interface{}{
		"id":          response.Id,
		"name":        response.Name,
		"goods_brief": response.GoodsBrief,
		"desc":        response.GoodsDesc,
		"ship_free":   response.ShipFree,
		"images":      response.Images,
		"desc_images": response.DescImages,
		"front_image": response.GoodsFrontImage,
		"shop_price":  response.ShopPrice,
		"ctegory": map[string]interface{}{
			"id":   response.Category.Id,
			"name": response.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   response.Brand.Id,
			"name": response.Brand.Name,
			"logo": response.Brand.Logo,
		},
		"is_hot":  response.IsHot,
		"is_new":  response.IsNew,
		"on_sale": response.OnSale,
	}
	ctx.JSON(http.StatusOK, goodsInfo)
}

func Delete(ctx *gin.Context) {
	goodsID := ctx.Param("id")
	goodsIDInt, err := strconv.Atoi(goodsID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}
	_, err = global.GoodsClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{Id: int32(goodsIDInt)})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.DeleteGoods failed", "msg", err.Error())
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func UpdateStatus(ctx *gin.Context) {
	goodsStatusForm := forms.GoodsStatusForm{}
	if err := ctx.ShouldBindJSON(&goodsStatusForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	goodsID := ctx.Param("id")
	goodsIDInt, err := strconv.Atoi(goodsID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	if _, err = global.GoodsClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:     int32(goodsIDInt),
		IsHot:  *goodsStatusForm.IsHot,
		IsNew:  *goodsStatusForm.IsNew,
		OnSale: *goodsStatusForm.OnSale,
	}); err != nil {
		zap.S().Errorw("global.GoodsClient.UpdateGoods failed", "msg", err.Error())
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
	})
}

func Update(ctx *gin.Context) {
	goodsForm := forms.GoodsForm{}
	if err := ctx.ShouldBindJSON(&goodsForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	goodsID := ctx.Param("id")
	goodsIDInt, err := strconv.Atoi(goodsID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	if _, err = global.GoodsClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              int32(goodsIDInt),
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.Brand,
	}); err != nil {
		zap.S().Errorw("global.GoodsClient.UpdateGoods failed", "msg", err.Error())
		HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

func HandleValidatorError(ctx *gin.Context, err error) {
	e, ok := err.(validator.ValidationErrors)
	if ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": removeTopStruct(e.Translate(global.Trans)),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": err.Error(),
	})
	return
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttp(ctx *gin.Context, err error) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": "服务不可用",
				})
			default:
				ctx.JSON(http.StatusNotFound, gin.H{
					"msg": "其他错误",
				})
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "未知错误",
			})
		}
	}
}
