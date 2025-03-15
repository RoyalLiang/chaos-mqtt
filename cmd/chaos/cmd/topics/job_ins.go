package topics

import (
	"fmt"
	"strconv"
	"strings"

	"fms-awesome-tools/cmd/chaos/internal/messages"
	tools "fms-awesome-tools/utils"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var JobInstructionCmd = &cobra.Command{
	Use:   "job_instruction",
	Short: "发送 job_instruction",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("job_instruction", "", generateTemplateParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateTemplateParam() string {
	dest := destination
	if strings.HasPrefix(destination, "PQC") {
		dest = "P," + destination + "          "
	}

	job := messages.JobInstruction{
		APMID: constants.VehicleID,
		Data: messages.JobInstructionData{
			ID: tools.GetVehicleTaskID(constants.VehicleID, dest, constants.Activity), RouteType: "G", RouteDAG: make([]messages.RouteDag, 0),
			Activity: constants.Activity, NextLocation: dest, NextLocationLane: lane, LiftType: liftSize, TargetDockPosition: targetDockPos,
			OperationalTypes: make([]string, 0),
			CNTRCategorys:    make([]string, 0), CNTRStatus: make([]string, 0), CNTRWeights: make([]float64, 0),
			CNTRNumbers: make([]string, 0), CNTRSizes: make([]string, 0), CNTRTypes: make([]string, 0),
			Cones: make([]int, 0), CNTRLocationsOnAPM: make([]string, 0), OperationalJobSequences: make([]string, 0),
			OperationalGroups: make([]string, 0), OperationalQCSequences: make([]string, 0), JobTypes: make([]string, 0),
			Urgents: make([]string, 0), DestLocations: make([]string, 0), DGGroups: make([]string, 0),
			DGS: make([]string, 0), ReferTemperatures: make([]float64, 0), IMOClass: make([]string, 0),
			OffloadSequences: make([]string, 0), TrailerPositions: make([]string, 0), WeightClass: make([]string, 0),
			PlugRequireds: make([]string, 0), SourceLocations: make([]string, 0), MotorDirections: make([]string, 0),
			AssignedCntrType: "GP", NumMountedCntr: 0, DualCycle: "N", AssignedCntrSize: "",
		},
	}
	switch constants.Activity {
	case 2, 3, 4:
		job.Data.AssignedCntrSize = strconv.FormatInt(container, 10)
	case 6, 7, 8:
		if liftSize == 1 || liftSize == 3 {
			job.Data.CNTRSizes = append(job.Data.CNTRSizes, strconv.FormatInt(container, 10))
		} else {
			job.Data.CNTRSizes = []string{"20", "20"}
		}

		if len(job.Data.CNTRSizes) == 1 {
			job.Data.CNTRLocationsOnAPM = []string{"1"}
		} else {
			job.Data.CNTRLocationsOnAPM = []string{"1", "3"}
		}
	}
	return job.String()
}

func init() {
	JobInstructionCmd.Flags().Int64VarP(&constants.Activity, "activity", "a", 1, activities)
	JobInstructionCmd.Flags().StringVarP(&destination, "destination", "d", "PQC921", "任务的目的地; QC: PQC921, Block: TB03_lane_2_slot_34")
	JobInstructionCmd.Flags().Int64VarP(&container, "container-size", "c", 40, "箱尺寸")
	JobInstructionCmd.Flags().StringVarP(&lane, "lane", "l", "2", "任务目的地车道")
	JobInstructionCmd.Flags().StringVarP(&targetDockPos, "dock-position", "x", "1", "任务目的点位; 1: 前箱, 3: 后箱, 5: 双20/单40")
	JobInstructionCmd.Flags().Int64VarP(&liftSize, "lift-size", "s", 1, "吊具尺寸; 1: 单20, 2: 双20, 3: 单40/45")
}
