package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/order_web/api/shopcart"
	"github.com/xlt/shop_web/order_web/middleware"
)

func InitShopCartRouter(v1Group *gin.RouterGroup) {
	shopCartRouter := v1Group.Group("/shop_cart").Use(middleware.JWTAuth())
	{
		shopCartRouter.GET("", shopcart.List)
		shopCartRouter.POST("", shopcart.New)
		shopCartRouter.DELETE("/:id", shopcart.Delete)
		shopCartRouter.PATCH("/:id", shopcart.Update)
	}

	zap.S().Infow("初始化订单购物车路由成功")
}
