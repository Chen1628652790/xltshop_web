package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xlt/shop_web/order_web/api/order"
	"github.com/xlt/shop_web/order_web/middleware"
	"go.uber.org/zap"
)

func InitOrderRouter(v1Group *gin.RouterGroup) {
	orderRouter := v1Group.Group("/orders")
	{
		orderRouter.GET("", middleware.JWTAuth(), order.List)
		orderRouter.POST("", middleware.JWTAuth(), middleware.AdminAuth(), order.New)
		orderRouter.GET("/:id", middleware.JWTAuth(), order.Detail)
	}

	zap.S().Infow("初始化订单路由成功")
}
