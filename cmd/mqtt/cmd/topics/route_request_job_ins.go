package topics

import (
	"fmt"
	"strings"

	"fms-awesome-tools/constants"
	tools "fms-awesome-tools/utils"

	"github.com/spf13/cobra"

	"fms-awesome-tools/cmd/mqtt/service"
)

const activities = "STANDBY = 1\nMOUNT = 2\nNO_YARD = 5\nOFFLOAD = 6\n"

var (
	destination   string
	container     string
	lane          int64
	targetDockPos int64
	liftSize      int64
)

var RouteJobCmd = &cobra.Command{
	Use:   "route_request_job_instruction",
	Short: "发送 route_request_job_instruction",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("route_request_job_instruction", constants.RouteRequestJobInstruction, generateTemplateParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateTemplateParam() interface{} {
	var dest = ""
	if strings.HasPrefix(destination, "PQC") {
		dest = "P," + destination + "          "
	}
	routeParam := constants.JobParam{
		ID:                 tools.GetVehicleTaskID(constants.VehicleID),
		VehicleID:          constants.VehicleID,
		Activity:           constants.Activity,
		Lane:               lane,
		Destination:        dest,
		LiftType:           liftSize,
		TargetDockPosition: targetDockPos,
	}
	return routeParam
}

func init() {
	RouteJobCmd.Flags().Int64VarP(&constants.Activity, "activity", "a", 1, activities)
	RouteJobCmd.Flags().StringVarP(&destination, "destination", "d", "PQC921", "任务的目的地; QC: PQC921, Block: TB03_lane_2_slot_34")
	RouteJobCmd.Flags().StringVarP(&container, "container-size", "c", "40", "箱尺寸")
	RouteJobCmd.Flags().Int64VarP(&lane, "lane", "l", 2, "任务目的地车道")
	RouteJobCmd.Flags().Int64VarP(&targetDockPos, "dock-position", "x", 1, "任务目的点位; 1: 前箱, 3: 后箱, 5: 双20/单40")
	RouteJobCmd.Flags().Int64VarP(&liftSize, "lift-size", "s", 1, "吊具尺寸; 1: 单20, 2: 双20, 3: 单40/45")
}
