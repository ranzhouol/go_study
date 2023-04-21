package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 设置默认值
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	fmt.Println(viper.Get("ContentDir"))
	fmt.Println(viper.Get("Taxonomies"))
}
