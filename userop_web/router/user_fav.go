package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xlt/shop_web/userop_web/api/user_fav"
	"github.com/xlt/shop_web/userop_web/middlewares"
)

func InitUserFavRouter(Router *gin.RouterGroup) {
	UserFavRouter := Router.Group("userfavs")
	{
		UserFavRouter.DELETE("/:id", middlewares.JWTAuth(), user_fav.Delete) // 删除收藏记录
		UserFavRouter.GET("/:id", middlewares.JWTAuth(), user_fav.Detail)    // 获取收藏记录
		UserFavRouter.POST("", middlewares.JWTAuth(), user_fav.New)          //新建收藏记录
		UserFavRouter.GET("", middlewares.JWTAuth(), user_fav.List)          //获取当前用户的收藏
	}
}
