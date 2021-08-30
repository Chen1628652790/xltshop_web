package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/xlt/shop_web/user_web/config"
)

var (
	ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator
)
