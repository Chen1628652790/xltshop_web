package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/goods_web/api/banner"
	"github.com/xlt/shop_web/goods_web/middleware"
)

func InitBannerRouter(v1Group *gin.RouterGroup) {
	bannerRouter := v1Group.Group("banners").Use(middleware.Trace())
	{
		bannerRouter.GET("", banner.List)                                                        // 轮播图列表页
		bannerRouter.DELETE("/:id", middleware.JWTAuth(), middleware.AdminAuth(), banner.Delete) // 删除轮播图
		bannerRouter.POST("", middleware.JWTAuth(), middleware.AdminAuth(), banner.New)          //新建轮播图
		bannerRouter.PUT("/:id", middleware.JWTAuth(), middleware.AdminAuth(), banner.Update)    //修改轮播图信息
	}

	zap.S().Infow("初始化轮播图路由成功")
}
