package main

import (
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"log"
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
	// jacket: leather
	// trousers: denim
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
	//	fmt.Println(viper.Get("name"))
	//	fmt.Println(viper.Get("clothing.jacket"))

	// 使用别名：修改任意一个，都会变化
	//viper.RegisterAlias("姓名", "name")
	//viper.Set("姓名", "LiMing")
	//fmt.Println(viper.Get("name"), viper.Get("姓名"))
	//viper.Set("name", "WangFang")
	//fmt.Println(viper.Get("name"), viper.Get("姓名"))

	// 使用环境变量
	//viper.AutomaticEnv()
	//viper.BindEnv()
	//viper.SetEnvPrefix()
	//viper.SetEnvKeyReplacer()
	//viper.AllowEmptyEnv()
	//viper.BindEnv("gopath", "GOPATH")
	//fmt.Println(viper.Get("gopath"))

	//// 使用flags
	//pflag.Int("password", 123456, "your password")
	//pflag.Parse()
	//viper.BindPFlags(pflag.CommandLine)
	//fmt.Println(viper.Get("password")) // 从viper而不是从pflag检索值
	//viper.Set("name", "dsd")
	//fmt.Println(viper.AllSettings(), viper.AllKeys())
	//fmt.Println(viper.IsSet("password"), viper.IsSet("name"))

	/*
		提取子树, 可用于实例化结构体
	*/
	//viper.SetConfigFile("./configuration/viper/config/test/subtree.yaml")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//sub1 := viper.Sub("app.cache1")
	//sub2 := viper.Sub("app.cache2")
	//fmt.Println(sub1.Get("max-items"))
	//fmt.Println(sub2.Get("max-items"))

	// 方法1
	//ca1 := NewCache(sub1)
	//ca2 := NewCache(sub2)

	// 方法2：反序列化
	//var ca1 cache
	//var ca2 cache
	//err = sub1.Unmarshal(&ca1)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//err = sub2.Unmarshal(&ca2)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(ca1.ItemVersion, ca2.ItemVersion)

	/*
		反序列化：将viper值解析到结构体、Map等
		Unmarshal(rawVal interface{})
		UnmarshalKey(key string, rawVal interface{})
	*/
	// 例1
	//type config struct {
	//	Port    int
	//	Name    string
	//	PathMap string `mapstructure:"path_map"`
	//}
	//
	//var C config
	//
	//viper.Set("port", 1)
	//viper.Set("name", "xx")
	//viper.Set("path_map", "/path1")
	//fmt.Println(viper.AllSettings())
	//err := viper.Unmarshal(&C)
	//if err != nil {
	//	fmt.Printf("unable to decode into struct, %v", err.Error())
	//}
	//fmt.Println(C.PathMap)

	// 例2：
	//v := viper.NewWithOptions(viper.KeyDelimiter(":")) // 修改键分隔符
	//v.SetDefault("chart:values", map[string]interface{}{
	//	"ingress": map[string]interface{}{
	//		"annotations": map[string]interface{}{
	//			"traefik.frontend.rule.type":                 "PathPrefix",
	//			"traefik.ingress.kubernetes.io/ssl-redirect": "true",
	//		},
	//	},
	//})
	//
	//type config struct {
	//	Chart struct {
	//		Values map[string]interface{}
	//	}
	//}
	//
	//var C config
	//
	//err := v.Unmarshal(&C) // 反序列化
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(C.Chart.Values)  // 由于结构体config只设置了键.Chart.Values，使用不了.Chart.Valuesy.ingress
	//v.WriteConfigAs("./xx.yaml") //保存

	/*
		序列化为字符串
	*/
	//v := viper.NewWithOptions(viper.KeyDelimiter(":")) // 修改键分隔符
	//v.SetDefault("chart:values", map[string]interface{}{
	//	"ingress": map[string]interface{}{
	//		"annotations": map[string]interface{}{
	//			"traefik.frontend.rule.type":                 "PathPrefix",
	//			"traefik.ingress.kubernetes.io/ssl-redirect": "true",
	//		},
	//	},
	//})
	//str1 := yamlStringSettings(v)
	//fmt.Println(str1)

	/*
		多个viper实例
		Viper是开箱即用的。你不需要配置或初始化即可开始使用Viper。
		你还可以在应用程序中创建许多不同的viper实例。每个都有自己独特的一组配置和值。
	*/
	//v1 := viper.New()
	//v2 := viper.New()
	//v1.SetDefault("ContentDir", "content1")
	//v2.SetDefault("ContentDir", "content2")
	//fmt.Println(v1.Get("contentdir"), v2.Get("contentdir"))
}

// 提取子树
type cache struct {
	MaxItems    int    `mapstructure:"max-items"`
	ItemSize    int    `mapstructure:"item-size"`
	ItemVersion string `mapstructure:"item-version"`
}

func NewCache(v *viper.Viper) cache {
	ca := cache{
		MaxItems:    v.GetInt("max-items"),
		ItemSize:    v.GetInt("item-size"),
		ItemVersion: v.GetString("item-version"),
	}
	return ca
}

// 序列化为字符串
func yamlStringSettings(viper *viper.Viper) string {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Printf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}
