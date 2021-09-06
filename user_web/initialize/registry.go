package initialize

import (
	"go.uber.org/zap"

	"github.com/xlt/shop_web/user_web/global"
	"github.com/xlt/shop_web/user_web/utils/register/consul"
)

func InitRegistry(serviceID string) {
	global.Registry = consul.NewRegistryClient(
		global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port,
	)
	if err := global.Registry.Register(
		global.ServerConfig.Host,
		global.ServerConfig.Port,
		global.ServerConfig.Name,
		global.ServerConfig.Tags,
		serviceID,
	); err != nil {
		zap.S().Errorw("服务注册失败", "msg", err.Error())
	}
}
