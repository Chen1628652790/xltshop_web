package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xlt/shop_web/order_web/api/shopcart"
	"github.com/xlt/shop_web/user_web/middleware"
	"go.uber.org/zap"
)

func InitShopCartRouter(v1Group *gin.RouterGroup) {
	shopCartRouter := v1Group.Group("/shopcarts").Use(middleware.JWTAuth())
	{
		shopCartRouter.GET("", shopcart.List)
		shopCartRouter.POST("", shopcart.New)
		shopCartRouter.DELETE("/:id", shopcart.Delete)
		shopCartRouter.PATCH("/:id", shopcart.Update)
	}

	zap.S().Infow("初始化订单购物车成功")
}
