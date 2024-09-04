package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/mqtt/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var VesselDepartCmd = &cobra.Command{
	Use:   "switch_mode",
	Short: "发送 switch_mode",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("switch_mode", constants.SwitchMode, generateSwitchModeParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}
