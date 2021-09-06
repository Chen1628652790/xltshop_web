package router

import (
	"github.com/gin-gonic/gin"

	"github.com/xlt/shop_web/goods_web/api/brand"
)

func InitBrandRouter(Router *gin.RouterGroup) {
	brandRouter := Router.Group("brands")
	{
		brandRouter.GET("", brand.BrandList)          // 品牌列表页
		brandRouter.DELETE("/:id", brand.DeleteBrand) // 删除品牌
		brandRouter.POST("", brand.NewBrand)          //新建品牌
		brandRouter.PUT("/:id", brand.UpdateBrand)    //修改品牌信息
	}

	categoryBrandRouter := Router.Group("categorybrands")
	{
		categoryBrandRouter.GET("", brand.CategoryBrandList)          // 类别品牌列表页
		categoryBrandRouter.DELETE("/:id", brand.DeleteCategoryBrand) // 删除类别品牌
		categoryBrandRouter.POST("", brand.NewCategoryBrand)          //新建类别品牌
		categoryBrandRouter.PUT("/:id", brand.UpdateCategoryBrand)    //修改类别品牌
		categoryBrandRouter.GET("/:id", brand.GetCategoryBrandList)   //获取分类的品牌
	}
}
