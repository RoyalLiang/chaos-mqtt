package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/mqtt/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var RouteRequestCmd = &cobra.Command{
	Use:   "route_request",
	Short: "发送 route_request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("route_request", constants.SwitchMode, generateSwitchModeParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}
