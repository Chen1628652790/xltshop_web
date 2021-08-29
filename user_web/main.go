package main

import (
	"go.uber.org/zap"

	"github.com/xlt/shop_web/user_web/initialize"
)

func main() {
	initialize.InitLogger()
	engine := initialize.InitRouter()

	zap.S().Infow("用户服务启动", "port", 8021)
	if err := engine.Run(":8021"); err != nil {
		zap.S().Fatal("用户服务启动失败")
	}

}
