package main

import (
	"fmt"
	"github.com/xlt/shop_web/user_web/global"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/user_web/initialize"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitTrans("zh")
	engine := initialize.InitRouter()

	zap.S().Infow("用户服务启动", "port", global.ServerConfig.Port)
	if err := engine.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Fatal("用户服务启动失败")
	}
}
