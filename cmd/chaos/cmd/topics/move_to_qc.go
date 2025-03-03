package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/internal/messages"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var MoveToQCCmd = &cobra.Command{
	Use:   "move_to_qc",
	Short: "发送 move_to_qc",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("move_to_qc", "", generateMoveToQC()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateMoveToQC() interface{} {
	return messages.MoveToQCRequest{
		APMID: constants.VehicleID,
		Data: messages.MOveToQCRequestData{
			RouteType: "G", RouteDag: make([]messages.RouteDag, 0),
		},
	}.String()
}
