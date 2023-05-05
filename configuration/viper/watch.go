package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

/*
	监控配置文件
*/

func main() {
	// 1. 读取
	viper.SetConfigFile("./configuration/viper/config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 2. 监控
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	for i := 0; i < 10; i++ {
		fmt.Println(i, viper.Get("username"))
		time.Sleep(2 * time.Second)
	}

}
