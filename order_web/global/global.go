package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/xlt/shop_web/order_web/config"
	"github.com/xlt/shop_web/order_web/proto"
	"github.com/xlt/shop_web/order_web/utils/register/consul"
)

var (
	// order_web服务配置
	ServerConfig = &config.ServerConfig{}

	// 翻译
	Trans ut.Translator

	// 订单服务(order_srv)连接
	OrderClient proto.OrderClient

	// 商品服务(goods_srv)连接
	GoodsClient proto.GoodsClient

	// 库存服务(inventory_srv)连接
	InventoryClient proto.InventoryClient

	// nacos配置
	NacosConfig = &config.NacosConfig{}

	// consul注册中心
	Registry consul.RegistryClient
)
