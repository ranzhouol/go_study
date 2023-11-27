package _case

import (
	"fmt"
	"time"
)

// 增删改查
func Crud() {
	temp := Teacher{
		Name:     "nick",
		Age:      40,
		Roles:    []string{"普通用户", "讲师"},
		Birthday: time.Now().Unix(),
		Salary:   12345.1234,
		Email:    "nick@voice.com",
	}
	t := temp

	res := DB.Create(&t)
	fmt.Println(res.RowsAffected, res.Error, t.ID)

	// 查找
	t1 := Teacher{}
	DB.First(&t1) //第一条记录
	fmt.Println(t1)

	// 更新
	t1.Name = "king"
	t1.Age = 31
	DB.Save(&t1)

	// 删除
	DB.Delete(&t1)
}
