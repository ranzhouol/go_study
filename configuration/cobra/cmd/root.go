package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 定义根命令 rootCmd
var rootCmd = &cobra.Command{
	// 定义命令参数名
	Use: "root",

	// 参数简介
	Short: "root命令的简介",

	// 参数详情
	Long: "root命令的详情介绍...",

	// 指定函数
	// Run func(cmd *Command, args []string)：只执行函数，
	// RunE func(cmd *Command, args []string) error: 带返回错误的执行函数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")

		// 打印标志的值
		fmt.Println(cmd.Flags().Lookup("viper").Value)
		fmt.Println(cmd.Flags().Lookup("author").Value)
		fmt.Println(cmd.Flags().Lookup("config").Value)
		fmt.Println(cmd.Flags().Lookup("license").Value)
		fmt.Println(cmd.Flags().Lookup("source").Value)
		// 打印viper中的值
		fmt.Println(viper.GetString("author"))
		fmt.Println(viper.GetString("license"))

		fmt.Println("root cmd run end")
	},
}

// 启动命令
func Execute() {
	rootCmd.Execute()
}

var cfgFile string
var userLicense string

func init() {
	// 调用自定义初始化函数
	cobra.OnInitialize(initConfig)

	// 为根命令添加bool类型和string类型的持久标志flags，即全局标志
	// 参数：flag名、默认值、说明
	rootCmd.PersistentFlags().Bool("viper", true, "viper的说明")
	rootCmd.PersistentFlags().String("test", "test", "test的说明")

	// 末尾带P的flags可以使用扩折号-的flag缩写
	// 参数：flag名、flag缩写、默认值、说明
	// 例如: -a
	rootCmd.PersistentFlags().StringP("author", "a", "RANZHOU", "")

	// 末尾带Var的flags可以使用一个相应的类型指针来保存flag值
	// 参数：指针、flag名、默认值、说明
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	// 也可以带缩写
	// 参数：指针、flag名、flag缩写、默认值、说明
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "")

	// 设置必填持久标志
	//rootCmd.MarkPersistentFlagRequired("config")
	// 设置必填本地标志
	//rootCmd.MarkFlagRequired("config")

	// 添加本地标志
	rootCmd.Flags().StringP("source", "s", "", "")

	// 将viper中的参数绑定到cobra命令
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))
	// 设置默认值
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")
}

// 加载配置文件
func initConfig() {
	if cfgFile == "" { //若没有指定配置文件
		return
	}

	// 读取配置文件
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		cobra.CheckErr(err)
	}

	fmt.Println("using config file:", viper.ConfigFileUsed())
}
