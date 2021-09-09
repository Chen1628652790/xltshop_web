package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xlt/shop_web/order_web/middleware"
	"github.com/xlt/shop_web/order_web/router"
)

func InitRouter() *gin.Engine {
	//todo 发布需要设置为release
	//gin.SetMode(global.ServerConfig.Mode)
	engine := gin.Default()
	engine.Use(middleware.Cors())

	engine.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	})

	v1Group := engine.Group("/o/v1")
	router.InitOrderRouter(v1Group)
	router.InitShopCartRouter(v1Group)

	return engine
}
