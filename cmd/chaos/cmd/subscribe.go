package cmd

import (
	"github.com/spf13/cobra"

	"fms-awesome-tools/cmd/chaos/service"
	tools "fms-awesome-tools/utils"
)

var name string

var subCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "订阅指定Topic的消息",
	Long:  tools.CustomTitle("订阅指定Topic的消息"),
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			_ = cmd.Help()
		} else {
			service.StartSubscribe(name)
		}
	},
}

func init() {
	subCmd.Flags().StringVarP(&name, "topic", "t", "", "topic名称🔠")
	rootCmd.AddCommand(subCmd)
}
