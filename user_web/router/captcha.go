package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/user_web/api"
)

func InitCaptchaRouter(v1Group *gin.RouterGroup) {
	captchaGroup := v1Group.Group("/base")

	captchaGroup.GET("/captcha", api.GetCaptcha)

	zap.S().Infow("初始化验证码路由成功")
}
