package main

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/xlt/shop_web/goods_web/global"
	"github.com/xlt/shop_web/goods_web/initialize"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitClient()
	initialize.InitTrans("zh")
	engine := initialize.InitRouter()

	if global.ServerConfig.Mode == "release" {
		// todo 线上部署需要随机端口
		//global.ServerConfig.Port = utils.GetFreePort()
	}

	zap.S().Infow("商品服务正在启动", "port", global.ServerConfig.Port)
	if err := engine.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Fatal("商品服务启动失败")
	}
}
