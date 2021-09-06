package global

import (
	ut "github.com/go-playground/universal-translator"

	"github.com/xlt/shop_web/user_web/config"
	"github.com/xlt/shop_web/user_web/proto"
	"github.com/xlt/shop_web/user_web/utils/register/consul"
)

var (
	ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator
	UserClient   proto.UserClient
	NacosConfig  = &config.NacosConfig{}
	Registry     consul.RegistryClient
)
