package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	id string
)

var VesselDepartCmd = &cobra.Command{
	Use:   "vessel_unberth",
	Short: "发送 vessel_unberth",
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" {
			fmt.Println("船舶ID不允许为空")
			_ = cmd.Help()
			return
		}
		if err := service.PublishAssignedTopic("vessel_unberth", constants.VesselUnberth, constants.VesselUnberthParam{VesselID: id}); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func init() {
	VesselDepartCmd.Flags().StringVarP(&id, "id", "i", "", "vessel info ID")
	_ = VesselBerthCmd.MarkFlagRequired("id")
}
