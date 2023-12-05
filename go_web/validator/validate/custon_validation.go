package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

/*
	3. 自定义验证
*/

type CustomValidateUser struct {
	Name   string `v:"required,is-nick"`
	Gender int    `v:"gender"`
}

func CustomValidate() {
	// 1. 注册自定义验证
	// tag名、验证函数
	validate.RegisterValidation("is-nick", ValidateIsNick)
	validate.RegisterValidation("gender", ValidateGender)

	// 2. 创建示例进行校验
	s := CustomValidateUser{
		Name:   "nick2",
		Gender: 1,
	}

	err := validate.Struct(s)
	if err != nil {
		fmt.Println(err)
	}
}

func ValidateIsNick(fl validator.FieldLevel) bool {
	return fl.Field().String() == "nick"
}

func ValidateGender(fl validator.FieldLevel) bool {
	i := fl.Field().Int()
	if i != 0 && i != 1 {
		return false
	}
	return true
}
