package topics

import (
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
		if err := service.PublishAssignedTopic("switch_mode", constants.SwitchMode, generateSwitchModeParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateSwitchModeParam() interface{} {
	return constants.SwitchModeParam{
		VehicleID: constants.VehicleID,
		Mode:      mode,
	}
}

func init() {
	SwitchCmd.Flags().StringVarP(&mode, "mode", "m", "OP", "设置单车模式; OP/TN/MA")
}
