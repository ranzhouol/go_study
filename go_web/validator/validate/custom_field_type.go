package validate

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
)

/*
	1. 自定义字段验证
*/

// 数据库用户, 使用自定义tag：v
type DbBackdUser struct {
	Name sql.NullString `v:"required"`
	Age  sql.NullInt64  `v:"required"`
	Loc  Location       `v:"required"`
}

// 自定义字段类型: 坐标
type Location struct {
	X float64
	Y float64
}

// 校验字段
func CustomField() {
	// 1. 注册校验方法以及需要校验的字段类型
	// 校验方法：ValidateValue
	// 后面的参数类型都是"database/sql"库中字段
	validate.RegisterCustomTypeFunc(ValidateValue, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

	// 校验方法：ValidateLocation
	// 校验字段：Location{}
	validate.RegisterCustomTypeFunc(ValidateLocation, Location{})

	// 2. 创建示例进行校验
	x := DbBackdUser{
		Name: sql.NullString{String: "nick", Valid: true},
		Age:  sql.NullInt64{Int64: 20, Valid: true},
		Loc:  Location{X: 10, Y: 10},
	}

	err := validate.Struct(x)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
	}
}

// 校验方法
func ValidateValue(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		} else {
			return nil
		}
	}
	return nil
}

// 自定义字段的校验方法
func ValidateLocation(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(Location); ok {
		if valuer.X != 0 && valuer.Y != 0 {
			return fmt.Sprintf("point(%v %v)", valuer.X, valuer.Y)
		}
	}

	return nil
}
