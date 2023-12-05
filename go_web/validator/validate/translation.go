package validate

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

/*
	4. 消息提示本地化
*/

var (
	// 翻译器
	uni *ut.UniversalTranslator
)

func Translation() {
	local := "en"

	// 英文包、中文包
	// 提供日期等通用内容的翻译
	en := en.New()
	zh := zh.New()

	// 初始化翻译器
	// 第一个参数为翻译失败时使用的翻译语种，后面的参数为支持的翻译语种
	uni = ut.New(en, zh, en)

	// 根据指定语种获取翻译器
	trans, _ := uni.GetTranslator(local)

	switch local {
	case "en":
		en_translations.RegisterDefaultTranslations(validate, trans)
	case "zh":
		zh_translations.RegisterDefaultTranslations(validate, trans)
	default:
		en_translations.RegisterDefaultTranslations(validate, trans)
	}

	// 使用第三方包默认翻译
	fmt.Println("默认翻译:--------------")
	translateAll(trans)
	// 可以重写翻译内容
	fmt.Println("自定义翻译:--------------")
	translateOverride(trans)
}

func translateAll(trans ut.Translator) {
	type User struct {
		UserName string `v:"required"`
		Tagline  string `v:"required,lt=10"`
		Tagline2 string `v:"required,gt=1"`
	}

	user := User{
		UserName: "Nick",
		Tagline:  "abcdefghijklmn",
		Tagline2: "1",
	}

	err := validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			msg := e.Translate(trans)
			fmt.Println(msg)
		}
	}
}

// 重写翻译
func translateOverride(trans ut.Translator) {
	local := trans.Locale()

	switch local {
	case "en":
		validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value! --override", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})
	case "zh":
		validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
			return ut.Add("required", "{0} 为必填字段! --override", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})
	}

	type User struct {
		UserName string `v:"required"`
	}

	u := User{}
	err := validate.Struct(u)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fmt.Println(e.Translate(trans))
		}
	}
}
