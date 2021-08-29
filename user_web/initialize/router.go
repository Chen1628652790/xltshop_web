package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/xlt/shop_web/user_web/router"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()

	v1Group := engine.Group("/v1")
	router.InitUserRouter(v1Group)

	return engine
}
