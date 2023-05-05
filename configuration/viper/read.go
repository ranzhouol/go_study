package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*
	读取配置文件、io.reader
*/

func main() {
	/*
		---
			【一】 配置文件
		---
	*/
	// 1. 设置默认值
	//viper.SetDefault("ContentDir", "content")
	//viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
	//
	//fmt.Println(viper.Get("ContentDir"))
	//fmt.Println(viper.Get("Taxonomies"))

	// 2. 读取配置文件，注意os.Getwd()获取的当前目录是 D:\Go\Go_WorkSpace\src\ranzhouol\go_study
	// 方法1：（最直接）直接按路径、名称、扩展名查找文件
	//viper.SetConfigFile("./configuration/viper/config/test/config.yaml") //指定配置文件路径，包括路径、名称、扩展名

	// 方法2：文件类型属于var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop",
	// "hcl", "tfvars", "dotenv", "env", "ini"}
	// 若SetConfigName不加扩展名，会根据文件名字ConfigName，依次从设置的第一个搜索路径ConfigPath中依次查找上述文件类型，
	// 找到第一个就返回
	//viper.SetConfigName("config")                        //指定配置文件名称，无扩展名（配置反而会出错），找到第一个就停止
	//viper.AddConfigPath("./configuration/viper/config/") //指定配置文件搜索路径，可以多次调用
	//viper.AddConfigPath("./configuration/viper/config/test/")

	// // 对于没有扩展名的文件，需要设置文件类型SetConfigType来进行解析
	//viper.SetConfigName("config") //指定配置文件名称
	//viper.SetConfigType("yaml")   //指定配置文件类型
	//viper.AddConfigPath("./configuration/viper/config/test1/")

	// 若SetConfigName加扩展名，则必须设置文件类型SetConfigType，否则找不到，依次从搜索路径中找到该文件
	//viper.SetConfigName("config.yaml")                   //指定配置文件名称，加扩展名，找到第一个就停止
	//viper.SetConfigType("yaml")                          //指定配置文件类型
	//viper.AddConfigPath("./configuration/viper/config/") //指定配置文件搜索路径，可以多次调用
	//viper.AddConfigPath("./configuration/viper/config/test/")

	//err := viper.ReadInConfig() //查找并读取配置文件
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(viper.Get("username"))

	/*
		---
			【二】 从io.Reader中读取配置
		---
	*/
	//	viper.SetConfigType("yaml")
	//
	//	var yamlExample = []byte(`
	//Hacker: true
	//name: steve
	//hobbies:
	//- skateboarding
	//- snowboarding
	//- go
	//clothing:
	//  jacket: leather
	//  trousers: denim
	//age: 35
	//eyes : brown
	//beard: true
	//`)
	//
	//	// bytes.NewBuffer: 缓冲byte类型的缓冲器
	//	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//fmt.Println(viper.Get("name"))

	// 使用别名
	//viper.RegisterAlias("姓名", "name")
	//viper.Set("name", "LiMing")
	//fmt.Println(viper.Get("name"))
	//fmt.Println(viper.Get("姓名"))

	// 使用环境变量
	//viper.AutomaticEnv()
	//viper.BindEnv()
	//viper.SetEnvPrefix()
	//viper.SetEnvKeyReplacer()
	//viper.AllowEmptyEnv()
	//viper.BindEnv("gopath", "GOPATH")
	//fmt.Println(viper.Get("gopath"))

	//
	pflag.Int("password", 123456, "your password")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}
