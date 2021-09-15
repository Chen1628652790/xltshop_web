package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/xlt/shop_web/goods_web/global"
	"github.com/xlt/shop_web/goods_web/proto"
	"github.com/xlt/shop_web/goods_web/utils/otgrpc"
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
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Errorw("grpc.Dial failed", "msg", err.Error())
	}

	global.GoodsClient = proto.NewGoodsClient(goodsConn)
}
