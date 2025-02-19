package cmd

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"
	tools "fms-awesome-tools/utils"

	"github.com/spf13/cobra"
)

var (
	start    bool
	dest     string
	lane     string
	auto     bool
	vehicles []string
)

var workflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "运行指定的workflow",
	Long:  tools.CustomTitle("运行指定的workflow"),
	Run: func(cmd *cobra.Command, args []string) {
		if start {
			startWorkflow()
		} else {
			_ = cmd.Help()
		}
	},
}

func startWorkflow() {
	if constants.Activity != 1 && constants.Activity != 2 && constants.Activity != 5 && constants.Activity != 6 {
		fmt.Printf("activity <%d> 不在可选范围内\n", constants.Activity)
		return
	}
	if (constants.Activity == 2 || constants.Activity == 6) && (dest == "" || lane == "") {
		fmt.Printf("activity 为 <%d> 时, destination与lane不可为空 \n", constants.Activity)
		return
	}

	if err := service.NewWorkflow(constants.Activity, lane, constants.VehicleID, dest, auto).StartWorkflow(); err != nil {
		fmt.Println("failed to start workflow:", err)
		return
	}
}

func init() {
	workflowCmd.Flags().BoolVarP(&start, "start", "s", false, "start workflow;")
	workflowCmd.Flags().Int64VarP(&constants.Activity, "activity", "a", 1, "STANDBY = 1\nMOUNT = 2\nNO_YARD = 5\nOFFLOAD = 6\n")
	workflowCmd.Flags().StringVarP(&constants.VehicleID, "truck", "v", "APM9001", "集卡号;")
	workflowCmd.Flags().StringVarP(&dest, "destination", "d", "", "任务的目的地; QC: PQC921, Block: Y,V,,TB01,32,32,10, ;")
	workflowCmd.Flags().StringVarP(&lane, "lane", "l", "2", "车道号;")
	workflowCmd.Flags().BoolVarP(&auto, "auto-callin", "", false, "是否自动发送call-in request(5s);")
	workflowCmd.Flags().StringSliceVarP(&vehicles, "vehicles", "", make([]string, 0), "")
	//workflowCmd.MarkFlagsRequiredTogether("truck", "activity")
	rootCmd.AddCommand(workflowCmd)
}
