package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/xlt/shop_web/user_web/middleware"
	"github.com/xlt/shop_web/user_web/router"
	"net/http"
)

func InitRouter() *gin.Engine {
	//todo 发布需要设置为release
	//gin.SetMode(global.ServerConfig.Mode)
	engine := gin.Default()
	engine.Use(middleware.Cors())

	engine.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})

	v1Group := engine.Group("/u/v1")
	router.InitUserRouter(v1Group)
	router.InitCaptchaRouter(v1Group)

	return engine
}
