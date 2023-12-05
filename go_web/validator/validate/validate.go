package validate

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	// 初始化校验器
	validate = validator.New()
	// 设置自定义tag: v
	validate.SetTagName("v")
}
