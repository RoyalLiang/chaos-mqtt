package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"fms-awesome-tools/cmd/chaos/service"
	tools "fms-awesome-tools/utils"
)

var (
	topic   string
	message string
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "发送指定Topic数据到配置的MQTT Server",
	Long:  tools.CustomTitle("发送指定Topic数据到配置的MQTT Server"),
	Run: func(cmd *cobra.Command, args []string) {
		if topic == "" && message == "" {
			_ = cmd.Help()
			return
		}

		mqtt, err := service.NewMQTTClientWithConfig("chaos")
		if err != nil {
			cobra.CheckErr(err)
		}

		if err := mqtt.Publish(topic, message); err != nil {
			fmt.Println("publish error:", err)
		}
	},
}

func init() {
	publishCmd.PersistentFlags().StringVarP(&topic, "topic", "t", "", "待发送的topic名称🔠")
	publishCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "待发送的topic消息体🔡")
	publishCmd.MarkFlagsRequiredTogether("topic", "message")
	rootCmd.AddCommand(publishCmd)
}
