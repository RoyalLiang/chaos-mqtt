package cmd

import (
	"fms-awesome-tools/cmd/chaos/service"

	"github.com/spf13/cobra"

	tools "fms-awesome-tools/utils"
)

var (
	name       string
	serverType string
)

var subCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "订阅指定Topic的消息",
	Long:  tools.CustomTitle("订阅指定Topic的消息"),
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			_ = cmd.Help()
		} else {
			if serverType == "mqtt" {
				service.StartSubscribe(name)
			} else if serverType == "redis" {
				service.StartRedisSubscribe(name)
			}
		}
	},
}

func init() {
	subCmd.Flags().StringVarP(&name, "topic", "t", "", "topic名称🔠")
	subCmd.Flags().StringVar(&serverType, "type", "mqtt", "订阅类型，mqtt/redis")
	rootCmd.AddCommand(subCmd)
}
