package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xlt/shop_web/goods_web/middleware"
	"github.com/xlt/shop_web/goods_web/router"
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

	v1Group := engine.Group("/g/v1")
	router.InitGoodsRouter(v1Group)
	router.InitCategoryRouter(v1Group)
	router.InitBannerRouter(v1Group)
	router.InitBrandRouter(v1Group)

	return engine
}
