package cmd

import (
	"fmt"

	"fms-awesome-tools/cmd/mqtt/service"

	"github.com/spf13/cobra"
)

var (
	assignedTopic string
)

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "publish assigned topic message to MQTT",
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
	rootCmd.AddCommand(topicCmd)
}
