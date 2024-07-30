package cmd

import (
	"fms-awesome-tools/cmd/mqtt/service"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	topic   string
	message string
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish a MQTT message to a topic",
	Run: func(cmd *cobra.Command, args []string) {
		if topic == "" && message == "" {
			cmd.Help()
			return
		}
		if err := service.MQTTClient.Publish(topic, message); err != nil {
			fmt.Println("publish error:", err)
		}
	},
}

func init() {
	publishCmd.PersistentFlags().StringVarP(&topic, "topic", "t", "", "待发送的topic名称")
	publishCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "待发送的topic消息体")
	publishCmd.MarkFlagsRequiredTogether("topic", "message")
	rootCmd.AddCommand(publishCmd)
}
