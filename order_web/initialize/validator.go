package initialize

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"

	"github.com/xlt/shop_web/order_web/global"
	regvalidator "github.com/xlt/shop_web/order_web/validator"
)

func InitTrans(locale string) {
	var err error

	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			//name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			name := fld.Tag.Get("json")
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			zap.S().Errorw("uni.GetTranslator failed", "msg", ok)
			return
		}

		switch locale {
		case "en":
			err = en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			err = zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			err = en_translations.RegisterDefaultTranslations(v, global.Trans)
		}
	}

	if err != nil {
		zap.S().Errorw("en_translations.RegisterDefaultTranslations failed", "msg", err.Error())
	}

	registerValidator()

	return
}

func registerValidator() {
	var err error

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("mobile", regvalidator.ValidateMobile)
		if err != nil {
			zap.S().Errorw("v.RegisterValidation failed", "msg", err.Error())
		}

		err = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0}非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
		if err != nil {
			zap.S().Errorw("v.RegisterTranslation failed", "msg", err.Error())
		}
	}
}
