package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/xlt/shop_web/goods_web/config"
	"github.com/xlt/shop_web/goods_web/proto"
	"github.com/xlt/shop_web/goods_web/utils/register/consul"
)

var (
	ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator
	//UserClient   proto.UserClient
	GoodsClient proto.GoodsClient
	NacosConfig = &config.NacosConfig{}
	Registry    consul.RegistryClient
)
