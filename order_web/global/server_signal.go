package global

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/xlt/shop_web/order_web/utils/register/consul"
)

func QuitServer(serviceID string) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := Registry.DeRegister(serviceID); err != nil {
		zap.S().Errorw("registerClient.DeRegister failed", "msg", err.Error())
		return
	}
	zap.S().Infow("注销服务成功", "port", ServerConfig.Port, "serviveID", serviceID)
}

func RegisterServer(serviceID string) {
	Registry = consul.NewRegistryClient(
		ServerConfig.ConsulInfo.Host,
		ServerConfig.ConsulInfo.Port,
	)
	if err := Registry.Register(
		ServerConfig.Host,
		ServerConfig.Port,
		ServerConfig.Name,
		ServerConfig.Tags,
		serviceID,
	); err != nil {
		zap.S().Errorw("服务注册失败", "msg", err.Error())
		return
	}
	zap.S().Infow("服务注册成功", "serviceID", serviceID)
}
