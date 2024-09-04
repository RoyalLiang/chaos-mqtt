package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/mqtt/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	manual int64
)

var CAllCmd = &cobra.Command{
	Use:   "call_in_request",
	Short: "发送 call_in_request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("call_in_request", constants.CallInRequest, generateCallInRequestParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateCallInRequestParam() interface{} {
	return constants.CallInRequestParam{
		VehicleID: constants.VehicleID,
		Mode:      manual,
	}
}

func init() {
	CAllCmd.Flags().Int64VarP(&manual, "manual", "m", 0, "优先级; 0: 正常, 1: 优先")
}
