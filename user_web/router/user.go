package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xlt/shop_web/user_web/api"
)

func InitUserRouter(v1Group *gin.RouterGroup) {
	userGroup := v1Group.Group("/user")

	userGroup.GET("/list", api.GetUserList)
}
