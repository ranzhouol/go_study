package _case

import (
	"fmt"
	"gorm.io/gorm"
)

// 事务开始前
func (t *Teacher) BeforeSave(tx *gorm.DB) error {
	fmt.Println("hook beforesave")
	return nil
}

// 事务提交后
func (t *Teacher) AfterSave(tx *gorm.DB) error {
	fmt.Println("hook aftersave")
	return nil
}

// 创建前
func (t *Teacher) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("hook BeforeCreate")
	return nil
}

// 创建后
func (t *Teacher) AfterCreate(tx *gorm.DB) error {
	fmt.Println("hook AfterCreate")
	return nil
}

// 更新前
func (t *Teacher) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println("hook BeforeUpdate")
	return nil
}

// 更新后
func (t *Teacher) AfterUpdate(tx *gorm.DB) error {
	fmt.Println("hook AfterUpdate")
	return nil
}

// 删除前
func (t *Teacher) BeforeDelete(tx *gorm.DB) error {
	fmt.Println("hook BeforeDelete")
	return nil
}

// 删除后
func (t *Teacher) AfterDelete(tx *gorm.DB) error {
	fmt.Println("hook AfterDelete")
	return nil
}

func (t *Teacher) AfterFind(tx *gorm.DB) error {
	fmt.Println("hook AfterFind")
	return nil
}
