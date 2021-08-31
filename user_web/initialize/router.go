package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/xlt/shop_web/user_web/middleware"
	"github.com/xlt/shop_web/user_web/router"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Cors())

	v1Group := engine.Group("/u/v1")
	router.InitUserRouter(v1Group)

	return engine
}
