package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xlt/shop_web/userop_web/api/message"
	"github.com/xlt/shop_web/userop_web/middlewares"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middlewares.JWTAuth())
	{
		MessageRouter.GET("", message.List) // 轮播图列表页
		MessageRouter.POST("", message.New) //新建轮播图
	}
}
