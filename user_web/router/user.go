package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/user_web/api"
)

func InitUserRouter(v1Group *gin.RouterGroup) {
	userGroup := v1Group.Group("/user")

	userGroup.GET("/list", api.GetUserList)
	userGroup.POST("/pwd_login", api.PasswordLogin)

	zap.S().Infow("初始化用户路由成功")
}
