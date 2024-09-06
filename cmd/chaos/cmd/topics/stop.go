package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	action int64
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "发送 stop",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("stop", constants.Stop, generateStopParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateStopParam() interface{} {
	return constants.StopParam{
		VehicleID: constants.VehicleID,
		Action:    action,
	}
}

func init() {
	StopCmd.Flags().Int64VarP(&action, "action", "a", 1, "stop action;")
}
