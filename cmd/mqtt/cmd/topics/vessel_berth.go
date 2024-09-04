package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/mqtt/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	vesselID  string
	direction string
	start     int64
	end       int64
	cranes    string
)

var VesselBerthCmd = &cobra.Command{
	Use:   "vessel_berth",
	Short: "发送 vessel_berth",
	Run: func(cmd *cobra.Command, args []string) {
		if vesselID == "" || cranes == "" || start == 0 || end == 0 {
			fmt.Println("参数缺失，请检查")
			return
		}

		if err := service.PublishAssignedTopic("vessel_berth", constants.VesselBerth, genVesselBerth()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func genVesselBerth() interface{} {
	return constants.VesselBerthParam{
		VesselID:  vesselID,
		Direction: direction,
		Cranes:    cranes,
		StartPos:  start,
		EndPos:    end,
	}
}

func init() {
	VesselBerthCmd.Flags().StringVarP(&vesselID, "id", "i", "", "vessel ID")
	VesselBerthCmd.Flags().StringVarP(&direction, "direction", "d", "P", "vessel direction; support S/P")
	VesselBerthCmd.Flags().StringVarP(&cranes, "cranes", "c", "", "vessel bind cranes; split by <,> between cranes")
	VesselBerthCmd.Flags().Int64VarP(&start, "start", "s", 0, "vessel start wharf mark")
	VesselBerthCmd.Flags().Int64VarP(&end, "end", "e", 0, "vessel end wharf mark")
	VesselBerthCmd.MarkFlagsRequiredTogether("id", "cranes", "start", "end")
}
