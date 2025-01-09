package cmd

import (
	"fms-awesome-tools/cmd/chaos/service"
	"github.com/spf13/cobra"
)

var name string

var subCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "订阅指定主题消息",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			cmd.Help()
		} else {
			service.StartSubscribe(name)
		}
	},
}

func init() {
	subCmd.Flags().StringVarP(&name, "topic", "t", "", "topic名称")
	rootCmd.AddCommand(subCmd)
}
