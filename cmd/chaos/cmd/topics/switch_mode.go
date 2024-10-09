package topics

import (
	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fmt"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	mode string
)

var SwitchCmd = &cobra.Command{
	Use:   "switch_mode",
	Short: "发送 switch_mode",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("switch_mode", "", generateSwitchModeParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateSwitchModeParam() interface{} {
	return messages.SwitchModeRequest{
		ApmId: constants.VehicleID,
		Data: messages.SwitchModeRequestData{
			SetMode: mode,
		},
	}.String()
}

func init() {
	SwitchCmd.Flags().StringVarP(&mode, "mode", "m", "OP", "设置单车模式; OP/TN/MA")
}
