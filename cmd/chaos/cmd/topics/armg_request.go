package topics

import (
	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	dist      float64
	armgCrane string
)

var ArmgCmd = &cobra.Command{
	Use:   "call_in_request",
	Short: "发送 call_in_request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("call_in_request", "", genArmgRequest()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func genArmgRequest() string {
	return messages.ArmgRequest{
		APMID: constants.VehicleID,
		Data: messages.ArmgRequestData{
			Valid:       1,
			DistRemain:  dist,
			CraneNumber: armgCrane,
		},
	}.String()
}

func init() {
	ArmgCmd.Flags().Float64VarP(&dist, "dist-remain", "d", 0, "CPS距离(毫米)")
	ArmgCmd.Flags().StringVarP(&armgCrane, "crane", "c", "", "目标RMG号")
}
