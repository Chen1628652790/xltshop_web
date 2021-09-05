package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/xlt/shop_web/goods_web/middleware"
	"github.com/xlt/shop_web/goods_web/router"
)

func InitRouter() *gin.Engine {
	//todo 发布需要设置为release
	//gin.SetMode(global.ServerConfig.Mode)
	engine := gin.Default()
	engine.Use(middleware.Cors())

	v1Group := engine.Group("/g/v1")
	router.InitGoodsRouter(v1Group)

	return engine
}
