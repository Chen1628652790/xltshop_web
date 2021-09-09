package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/xlt/shop_web/order_web/global"
	"github.com/xlt/shop_web/order_web/proto"
)

func InitClient() {
	initOrderClientLoadBalance()
}

func initOrderClientLoadBalance() {
	orderConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port,
			global.ServerConfig.OrderSrvInfo.Name,
		),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("grpc.Dial failed", "msg", err.Error())
	}

	global.OrderClient = proto.NewOrderClient(orderConn)
	zap.S().Infow("连接OrderServer成功", "name", global.ServerConfig.OrderSrvInfo.Name)
}
