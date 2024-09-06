package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/service"
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
	VesselBerthCmd.Flags().StringVarP(&vesselID, "id", "i", "", "船舶ID")
	VesselBerthCmd.Flags().StringVarP(&direction, "direction", "d", "P", "船舶停靠方向 S(右舷)/P(左舷)")
	VesselBerthCmd.Flags().StringVarP(&cranes, "cranes", "c", "", "船舶作业岸桥列表(使用,号分隔)")
	//VesselBerthCmd.Flags().StringP("cranes", "c", "", "船舶作业岸桥列表()")
	VesselBerthCmd.Flags().Int64VarP(&start, "start", "s", 0, "船舶起始wharf mark")
	VesselBerthCmd.Flags().Int64VarP(&end, "end", "e", 0, "船舶终止wharf mark")
	VesselBerthCmd.MarkFlagsRequiredTogether("id", "cranes", "start", "end")
}
