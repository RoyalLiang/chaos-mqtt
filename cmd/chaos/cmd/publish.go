package cmd

import (
	tools "fms-awesome-tools/utils"
	"github.com/spf13/cobra"
)

var (
	topic   string
	message string
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish the message of assigned topic to MQTT",
	Long:  tools.CustomTitle("publish the message of assigned topic to MQTT"),
	Run: func(cmd *cobra.Command, args []string) {
		if topic == "" && message == "" {
			_ = cmd.Help()
			return
		}
		//if err := service.MQTTClient.Publish(topic, message); err != nil {
		//	fmt.Println("publish error:", err)
		//}
	},
}

func init() {
	publishCmd.PersistentFlags().StringVarP(&topic, "topic", "t", "", "待发送的topic名称")
	publishCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "待发送的topic消息体")
	publishCmd.MarkFlagsRequiredTogether("topic", "message")
	rootCmd.AddCommand(publishCmd)
}
