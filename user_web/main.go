package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/xlt/shop_web/user_web/global"
	"github.com/xlt/shop_web/user_web/initialize"
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

	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	initialize.InitRegistry(serviceID)

	zap.S().Infow("用户服务启动", "port", global.ServerConfig.Port)
	go func() {
		if err := engine.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Fatal("用户服务启动失败")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := global.Registry.DeRegister(serviceID); err != nil {
		zap.S().Errorw("global.Registry.DeRegister failed", "msg", err.Error())
		return
	}
	zap.S().Infow("注销服务成功", "port", global.ServerConfig.Port, "serviceID", serviceID)
}
