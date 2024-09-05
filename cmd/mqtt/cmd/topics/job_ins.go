package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/mqtt/service"
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

func init() {
	RouteJobCmd.Flags().Int64VarP(&constants.Activity, "job-activity", "", 1, activities)
	RouteJobCmd.Flags().StringVarP(&destination, "job-destination", "", "PQC921", "任务的目的地; QC: PQC921, Block: TB03_lane_2_slot_34")
	RouteJobCmd.Flags().StringVarP(&container, "job-container-size", "", "40", "箱尺寸")
	RouteJobCmd.Flags().Int64VarP(&lane, "job-lane", "", 2, "任务目的地车道")
	RouteJobCmd.Flags().Int64VarP(&targetDockPos, "job-dock-position", "", 1, "任务目的点位; 1: 前箱, 3: 后箱, 5: 双20/单40")
	RouteJobCmd.Flags().Int64VarP(&liftSize, "job-lift-size", "", 1, "吊具尺寸; 1: 单20, 2: 双20, 3: 单40/45")
}
