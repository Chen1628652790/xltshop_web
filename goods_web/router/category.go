package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/goods_web/api/category"
)

func InitCategoryRouter(v1Group *gin.RouterGroup) {
	categoryRouter := v1Group.Group("categorys")
	{
		categoryRouter.GET("", category.List)          // 商品类别列表页
		categoryRouter.DELETE("/:id", category.Delete) // 删除分类
		categoryRouter.GET("/:id", category.Detail)    // 获取分类详情
		categoryRouter.POST("", category.New)          //新建分类
		categoryRouter.PUT("/:id", category.Update)    //修改分类信息
	}

	zap.S().Infow("初始化分类路由成功")
}
