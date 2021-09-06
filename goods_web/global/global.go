package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/xlt/shop_web/goods_web/config"
	"github.com/xlt/shop_web/goods_web/proto"
	"github.com/xlt/shop_web/goods_web/utils/register/consul"
)

var (
	// goods_web服务配置
	ServerConfig = &config.ServerConfig{}

	// 翻译
	Trans ut.Translator

	// 商品服务(goods_srv)链接
	GoodsClient proto.GoodsClient

	// nacos配置
	NacosConfig = &config.NacosConfig{}

	// consul注册中心
	Registry consul.RegistryClient
)
