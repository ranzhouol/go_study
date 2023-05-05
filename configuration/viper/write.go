package main

import (
	"fmt"
	"github.com/spf13/viper"
)

/*
	写入配置文件
	WriteConfig - 将当前的viper配置写入预定义的路径并覆盖（如果存在的话）。如果没有预定义的路径，则报错。
	SafeWriteConfig - 将当前的viper配置写入预定义的路径。如果没有预定义的路径，则报错。如果存在，将不会覆盖当前的配置文件。
	WriteConfigAs - 将当前的viper配置写入给定的文件路径。将覆盖给定的文件(如果它存在的话)。
	SafeWriteConfigAs - 将当前的viper配置写入给定的文件路径。不会覆盖给定的文件(如果它存在的话)。
*/

func main() {
	// 1. 设置默认值或读取配置文件
	viper.SetDefault("UserName2", "testuser")
	viper.SetDefault("Message", map[string]string{"sex": "male", "age": "18"})

	viper.SetConfigFile("./configuration/viper/config/config.yaml")
	// 若只设置了配置文件路径，没有读取文件内容，执行写入WriteConfig操作会将文件置空，但SafeWriteConfig()不会
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(viper.Get("username"))
	// 2. 修改值
	viper.Set("username", "config-yaml-2")
	// 添加新配置
	viper.Set("sex", "male")
	fmt.Println(viper.Get("username"))

	// 3. 写入
	err = viper.WriteConfig() // 将当前的配置写入文件
	if err != nil {
		fmt.Println(err.Error())
	}

	// 指定文件名
	//err = viper.WriteConfigAs("./configuration/viper/config/testconfig.json") // 指定保存文件名
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
}
