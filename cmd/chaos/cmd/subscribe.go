package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var name string

var subCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "订阅指定主题消息",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Print("未指定topic名称, 订阅失败...")
		} else {

		}
	},
}

func init() {
	subCmd.Flags().StringVarP(&name, "topic", "t", "", "topic名称")
	rootCmd.AddCommand(topicCmd)
}
