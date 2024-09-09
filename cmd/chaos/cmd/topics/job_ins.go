package topics

import (
	tools "fms-awesome-tools/utils"
	"fmt"
	"strings"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var JobInstructionCmd = &cobra.Command{
	Use:   "job_instruction",
	Short: "发送 job_instruction",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("job_instruction", constants.JobInstruction, generateTemplateParam()); err != nil {
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
	JobInstructionCmd.Flags().Int64VarP(&constants.Activity, "activity", "a", 1, activities)
	JobInstructionCmd.Flags().StringVarP(&destination, "destination", "d", "PQC921", "任务的目的地; QC: PQC921, Block: TB03_lane_2_slot_34")
	JobInstructionCmd.Flags().Int64VarP(&container, "container-size", "c", 40, "箱尺寸")
	JobInstructionCmd.Flags().StringVarP(&lane, "lane", "l", "2", "任务目的地车道")
	JobInstructionCmd.Flags().StringVarP(&targetDockPos, "dock-position", "x", "1", "任务目的点位; 1: 前箱, 3: 后箱, 5: 双20/单40")
	JobInstructionCmd.Flags().Int64VarP(&liftSize, "lift-size", "s", 1, "吊具尺寸; 1: 单20, 2: 双20, 3: 单40/45")
}
