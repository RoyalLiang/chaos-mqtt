package topics

import (
	tools "fms-awesome-tools/utils"
	"fmt"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	dockActivity int64
	location     int64
)

var DockPositionCmd = &cobra.Command{
	Use:   "dock_position",
	Short: "发送 dock_position",
	Long:  tools.CustomTitle("发送 dock_position"),
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("dock_position", constants.DockPosition, generateDockPositionParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateDockPositionParam() interface{} {
	return constants.DockPositionParam{
		VehicleID:   constants.VehicleID,
		Activity:    dockActivity,
		ConLocation: location,
	}
}

func init() {
	DockPositionCmd.Flags().Int64VarP(&dockActivity, "activity", "a", 1, "activity")
	DockPositionCmd.Flags().Int64VarP(&location, "container-location", "p", 1, "cntr location on apm")
}
