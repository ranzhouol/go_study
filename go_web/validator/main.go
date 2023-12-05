package main

import "ranzhouol/go_study/go_web/validator/validate"

func main() {
	// 1. 自定义字段验证
	//validate.CustomField()

	// 2. 结构体补充验证
	//validate.StructLevel()

	// 3. 自定义验证
	//validate.CustomValidate()

	// 4. 消息提示本地化
	validate.Translation()
}
