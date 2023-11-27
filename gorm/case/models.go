package _case

import "gorm.io/gorm"

func init() {
	DB.Migrator().AutoMigrate(Teacher{})
}

// gorm model
type Roles []string
type Teacher struct {
	gorm.Model
	Name     string  `gorm:"size:256"`
	Email    string  `gorm:"size:256"`
	Salary   float64 `gorm:"scale:2;precision:7"`           //小数位数，列宽
	Age      uint8   `gorm:"check:age>30"`                  //check约束
	Birthday int64   `gorm:"serializer:unixtime;type:time"` //序列化, 转换为时间戳类型
	Roles    Roles   `gorm:"serializer:json"`               //序列化, 转换为json类型
}
