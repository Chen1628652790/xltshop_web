package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/satori/go.uuid"
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

	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	initialize.InitRegistry(serviceId)

	zap.S().Infow("商品服务正在启动", "port", global.ServerConfig.Port)
	go func() {
		if err := engine.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Fatal("商品服务启动失败")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := global.Registry.DeRegister(serviceId); err != nil {
		zap.S().Errorw("registerClient.DeRegister failed", "msg", err.Error())
		return
	}
	zap.S().Infow("注销服务成功", "port", global.ServerConfig.Port, "serviveID", serviceId)
}
