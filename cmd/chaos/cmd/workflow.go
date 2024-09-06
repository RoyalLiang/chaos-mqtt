package cmd

import (
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"
	tools "fms-awesome-tools/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	start     bool
	vehicleID string
	dest      string
	lane      int64
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
	if (constants.Activity == 2 || constants.Activity == 6) && (dest == "" || lane == 0) {
		fmt.Printf("activity 为 <%d> 时, destination与lane不可为空 \n", constants.Activity)
		return
	}

	if err := service.StartWorkflow(); err != nil {
		fmt.Println("failed to start workflow:", err)
		return
	}

	//dest := tools.ParseDestination(dest)
	//workflow := service.Workflow{
	//	UUID:     tools.GenerateUUID(),
	//	Truck:    vehicleID,
	//	Activity: constants.Activity,
	//}
}

func init() {
	workflowCmd.Flags().BoolVarP(&start, "start", "s", false, "start workflow")
	workflowCmd.Flags().Int64VarP(&constants.Activity, "activity", "a", 1, "STANDBY = 1\nMOUNT = 2\nNO_YARD = 5\nOFFLOAD = 6\n")
	workflowCmd.Flags().StringVarP(&vehicleID, "truck", "v", "APM9001", "集卡号")
	workflowCmd.Flags().StringVarP(&dest, "destination", "d", "", "目的地")
	workflowCmd.Flags().Int64VarP(&lane, "lane", "l", 0, "车道号")
	workflowCmd.MarkFlagsRequiredTogether("truck", "activity")
	rootCmd.AddCommand(workflowCmd)
}
