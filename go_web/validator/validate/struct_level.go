package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

/*
	2. 结构体补充验证
*/

type StructLevelUser struct {
	FirstName string
	LastName  string
	Age       uint8  `v:"gte=0,lte=130"` //0<=Age<=130, gte大于等于，lte小于等于，可以直接用
	Email     string `fld:"e-mail" v:"required,email"`
}

// 结构体补充验证
func StructLevel() {
	// 注册tag name，可以修改tag反射对应的字段名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		// 获取tag:fld对应的值
		name := strings.SplitN(field.Tag.Get("fld"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// 1. 注册结构体验证
	// 结构体补充验证方法：UserStructLevelValidation
	// 结构体：StructLevelUser{}
	validate.RegisterStructValidation(UserStructLevelValidation, StructLevelUser{})

	// 2. 创建示例进行校验
	user := &StructLevelUser{
		FirstName: "ran",
		LastName:  "zhou",
		Age:       100,
		Email:     "nick@voice.com",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
	}
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(StructLevelUser)
	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		// 括号中的参数: 结构体中的字段、反射后的字段名、结构体字段名、tag名、其它参数
		sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
		sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
	}
}
