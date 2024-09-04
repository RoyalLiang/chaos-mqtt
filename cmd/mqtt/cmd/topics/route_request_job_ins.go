package topics

import (
	"fms-awesome-tools/constants"
	"fmt"

	"github.com/spf13/cobra"

	"fms-awesome-tools/cmd/mqtt/service"
)

const activities = "STANDBY = 1\nMOUNT = 2\nNO_YARD = 5\nOFFLOAD = 6\n"

var (
	activity     int64
	nextLocation string
	nextLane     int64
)

var RouteJobCmd = &cobra.Command{
	Use:   "route_request_job_instruction",
	Short: "发送 route_request_job_instruction",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishRouteRequestJobInstruction("route_request_job_instruction", constants.VehicleID, nextLocation, nextLane, activity); err != nil {
			fmt.Println("error to publish: ", err)
		}
	},
}

func init() {
	RouteJobCmd.Flags().Int64VarP(&activity, "activity", "a", 1, activities)
	RouteJobCmd.Flags().StringVarP(&nextLocation, "destination", "d", "", "任务的目的地; QC: PQC921, Block: TB03_lane_2_slot_34")
	RouteJobCmd.Flags().Int64VarP(&nextLane, "lane", "l", 1, "任务目的地车道")
}
