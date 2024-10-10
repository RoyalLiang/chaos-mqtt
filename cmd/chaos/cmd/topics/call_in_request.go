package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	manual int64
	crane  string
)

var CAllCmd = &cobra.Command{
	Use:   "call_in_request",
	Short: "发送 call_in_request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("call_in_request", "", genCallInRequest()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func genCallInRequest() string {
	return messages.CallInRequest{
		APMID: constants.VehicleID,
		Data: messages.CallInRequestData{
			Crane:      crane,
			CallInMode: manual,
		},
	}.String()
}

func init() {
	CAllCmd.Flags().Int64VarP(&manual, "manual", "m", 0, "优先级; 0: 正常, 1: 优先")
	CAllCmd.Flags().StringVarP(&crane, "crane", "c", "", "目的QC")
}
