package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var level2Cmd = &cobra.Command{
	Use: "l2",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("l2 cmd run begin")

		// 打印标志的值
		fmt.Println(cmd.Flags().Lookup("viper").Value)
		fmt.Println(cmd.Flags().Lookup("author").Value)
		fmt.Println(cmd.Flags().Lookup("config").Value)
		fmt.Println(cmd.Flags().Lookup("license").Value)
		fmt.Println(viper.GetString("author"))
		fmt.Println(viper.GetString("license"))

		fmt.Println("l2 cmd run end")
	},
}

var curArgsCheck = &cobra.Command{
	Use: "cuscheck",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cuscheck cmd run start")
		fmt.Println(args)
		fmt.Println("cuscheck cmd run end")
	},
	// 自定义参数验证
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少输入1个参数")
		} else if len(args) > 2 {
			return errors.New("最多输入2个参数")
		}

		return nil
	},
}

var onlyArgsCmd = &cobra.Command{
	Use:       "only",
	ValidArgs: []string{"123", "abc", "rz"},
	Args:      cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("only cmd run start")
		fmt.Println(args)
		fmt.Println("only cmd run end")
	},
}

var rangeArgsCmd = &cobra.Command{
	Use:  "range",
	Args: cobra.RangeArgs(2, 4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("range cmd run start")
		fmt.Println(args)
		fmt.Println("range cmd run end")
	},
}

func init() {
	// 为init命令添加l2子命令
	initCmd.AddCommand(level2Cmd)

	// 为root命令添加子命令
	rootCmd.AddCommand(curArgsCheck)
	rootCmd.AddCommand(onlyArgsCmd)
	rootCmd.AddCommand(rangeArgsCmd)
}
