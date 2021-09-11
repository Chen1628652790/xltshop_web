package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/xlt/shop_web/userop_web/config"
	"github.com/xlt/shop_web/userop_web/proto"
)

var (
	Trans ut.Translator

	ServerConfig = &config.ServerConfig{}

	NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient

	MessageClient proto.MessageClient
	AddressClient proto.AddressClient
	UserFavClient proto.UserFavClient
)
