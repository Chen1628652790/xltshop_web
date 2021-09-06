package banner

import (
	"context"
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
	rsp, err := global.GoodsClient.BannerList(context.Background(), &empty.Empty{})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.BannerList failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["index"] = value.Index
		reMap["image"] = value.Image
		reMap["url"] = value.Url

		result = append(result, reMap)
	}

	ctx.JSON(http.StatusOK, result)
}

func New(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	rsp, err := global.GoodsClient.CreateBanner(context.Background(), &proto.BannerRequest{
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
		Image: bannerForm.Image,
	})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.CreateBanner failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	response := make(map[string]interface{})
	response["id"] = rsp.Id
	response["index"] = rsp.Index
	response["url"] = rsp.Url
	response["image"] = rsp.Image

	ctx.JSON(http.StatusOK, response)
}

func Update(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		zap.S().Errorw("ctx.ShouldBindJSON failed", "msg", err.Error())
		return
	}

	bannerID := ctx.Param("id")
	bannerIDInt, err := strconv.Atoi(bannerID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}

	_, err = global.GoodsClient.UpdateBanner(context.Background(), &proto.BannerRequest{
		Id:    int32(bannerIDInt),
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
	})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.UpdateBanner failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func Delete(ctx *gin.Context) {
	bannerID := ctx.Param("id")
	bannerIDInt, err := strconv.Atoi(bannerID)
	if err != nil {
		zap.S().Errorw("strconv.Atoi failed", "msg", err.Error())
		return
	}
	_, err = global.GoodsClient.DeleteBanner(context.Background(), &proto.BannerRequest{Id: int32(bannerIDInt)})
	if err != nil {
		zap.S().Errorw("global.GoodsClient.DeleteBanner failed", "msg", err.Error())
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, "")
}
