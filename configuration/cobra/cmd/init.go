package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"create"},
	Short:   "init short",
	Long:    "init long",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init cmd run begin")

		// 打印标志的值
		fmt.Println(cmd.Flags().Lookup("viper").Value)
		fmt.Println(cmd.Flags().Lookup("author").Value)
		fmt.Println(cmd.Flags().Lookup("config").Value)
		fmt.Println(cmd.Flags().Lookup("license").Value)
		fmt.Println(cmd.Flags().Lookup("initsource").Value)
		// 不能使用source标志flag，因为是root命令的本地标志
		//fmt.Println(cmd.Flags().Lookup("source").Value)
		// 打印viper中的值
		fmt.Println(viper.GetString("author"))
		fmt.Println(viper.GetString("license"))

		fmt.Println("init cmd run end")
	},
}

func init() {
	// 为init命令添加标志
	initCmd.Flags().StringP("initsource", "i", "", "")
	// 将标志设置为必填
	initCmd.MarkFlagRequired("initsource")

	// 为根命令添加子命令init
	rootCmd.AddCommand(initCmd)
}
