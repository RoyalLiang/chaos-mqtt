package cmd

import (
	tools "fms-awesome-tools/utils"
	"fmt"

	"fms-awesome-tools/cmd/chaos/cmd/topics"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	assignedTopic string
	listTopics    bool
)

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "publish assigned topic message to MQTT",
	Long:  tools.CustomTitle("publish assigned topic message to MQTT"),
	Run: func(cmd *cobra.Command, args []string) {
		if listTopics {
			listAvcsTopics()
		} else if assignedTopic != "" {
			fmt.Println("assigned topic message to ", assignedTopic)
		} else {
			cmd.Help()
		}
	},
}

func listAvcsTopics() {
	fmt.Println("=========================================================")
	for _, v := range constants.TopicFromAVCS {
		fmt.Println(v)
	}
	fmt.Println("=========================================================")
}

func init() {
	topicCmd.PersistentFlags().StringVarP(&constants.VehicleID, "vehicle", "v", "APM9001", "集卡号")

	topicCmd.Flags().StringVarP(&assignedTopic, "topic", "t", "", "指定的消息topic")
	topicCmd.Flags().BoolVarP(&listTopics, "list", "l", false, "列出AVCS的 topic 列表")
	topicCmd.MarkFlagsMutuallyExclusive("topic", "list")

	topicCmd.AddCommand(topics.CAllCmd)
	topicCmd.AddCommand(topics.RouteJobCmd)
	topicCmd.AddCommand(topics.JobInstructionCmd)
	topicCmd.AddCommand(topics.SwitchCmd)
	topicCmd.AddCommand(topics.DockPositionCmd)
	topicCmd.AddCommand(topics.IngressToCallInCmd)
	topicCmd.AddCommand(topics.MoveToQCCmd)
	topicCmd.AddCommand(topics.VesselBerthCmd)
	topicCmd.AddCommand(topics.VesselDepartCmd)
	topicCmd.AddCommand(topics.RouteRequestCmd)
	topicCmd.AddCommand(topics.LogonResponseCmd)
	topicCmd.AddCommand(topics.LogonRequestCmd)
	topicCmd.AddCommand(topics.StopCmd)
	topicCmd.AddCommand(topics.FunctionalCmd)
	topicCmd.AddCommand(topics.ArmgCmd)
	rootCmd.AddCommand(topicCmd)
}
