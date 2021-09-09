package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xlt/shop_web/goods_web/middleware"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/goods_web/api/goods"
)

func InitGoodsRouter(v1Group *gin.RouterGroup) {
	goodsRouter := v1Group.Group("/goods")
	{
		goodsRouter.GET("/", goods.List)
		goodsRouter.POST("/", middleware.JWTAuth(), middleware.AdminAuth(), goods.New)
		goodsRouter.GET("/:id", goods.Detail)
		goodsRouter.DELETE("/:id", goods.Delete)
		goodsRouter.PATCH("/:id", middleware.JWTAuth(), middleware.AdminAuth(), goods.UpdateStatus)
		goodsRouter.PUT("/:id", middleware.JWTAuth(), middleware.AdminAuth(), goods.Update)
	}

	zap.S().Infow("初始化商品路由成功")
}
