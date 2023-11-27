package _case

import (
	"gorm.io/gorm"
	"time"
)

// 会话
func Session() {
	tx := DB.Session(&gorm.Session{
		PrepareStmt:              true, //预编译
		SkipHooks:                true, //跳过钩子
		DisableNestedTransaction: true, //禁用嵌套事务
	})

	temp := Teacher{
		Name:     "nick",
		Age:      40,
		Roles:    []string{"普通用户", "讲师"},
		Birthday: time.Now().Unix(),
		Salary:   12345.1234,
		Email:    "nick@voice.com",
	}
	t := temp
	tx.Create(&t)
}
