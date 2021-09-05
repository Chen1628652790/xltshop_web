package initialize

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/xlt/shop_web/goods_web/global"
	"github.com/xlt/shop_web/goods_web/proto"
)

func InitClient() {
	initGoodsClientLoadBalance()
}

func initGoodsClientLoadBalance() {
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port,
			global.ServerConfig.GoodsSrvInfo.Name,
		),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorw("grpc.Dial failed", "msg", err.Error())
	}

	global.GoodsClient = proto.NewGoodsClient(goodsConn)
}

func initUserClient() {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d",
		global.ServerConfig.ConsulInfo.Host,
		global.ServerConfig.ConsulInfo.Port,
	)

	consulClient, err := api.NewClient(cfg)
	if err != nil {
		zap.S().Errorw("api.NewClient failed", "msg", err.Error())
		return
	}

	data, err := consulClient.Agent().ServicesWithFilter(
		fmt.Sprintf(`Service == "%s"`, global.ServerConfig.GoodsSrvInfo.Name),
	)
	if err != nil {
		zap.S().Errorw("consulClient.Agent().ServicesWithFilter failed", "msg", err.Error())
		return
	}

	var userSrvHost string
	var userSrvPort int
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("grpc.Dial failed", "msg", err.Error())
		return
	}

	// todo 负载均衡、健康检查、服务地址更改、连接池
	global.GoodsClient = proto.NewGoodsClient(conn)
	zap.S().Infow("proto.NewUserClient success", "msg", "连接用户服务成功")
}
