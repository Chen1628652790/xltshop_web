package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/goods_web/api/goods"
)

func InitGoodsRouter(v1Group *gin.RouterGroup) {
	goodsGroup := v1Group.Group("/goods")

	goodsGroup.GET("", goods.List)

	zap.S().Infow("初始化商品路由成功")
}
