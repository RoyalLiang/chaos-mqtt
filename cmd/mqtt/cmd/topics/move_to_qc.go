package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/mqtt/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var MoveToQCCmd = &cobra.Command{
	Use:   "move_to_qc",
	Short: "发送 move_to_qc",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("move_to_qc", constants.MoveToQC, generateMoveToQC()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateMoveToQC() interface{} {
	return constants.IngressToCallInParam{
		VehicleID: constants.VehicleID,
	}
}
