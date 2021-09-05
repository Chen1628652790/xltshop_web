package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/goods_web/api/goods"
	"github.com/xlt/shop_web/goods_web/middleware"
)

func InitGoodsRouter(v1Group *gin.RouterGroup) {
	goodsGroup := v1Group.Group("/user")

	goodsGroup.GET("/list", middleware.JWTAuth(), middleware.AdminAuth(), goods.GetGoodsList)

	zap.S().Infow("初始化用户路由成功")
}
