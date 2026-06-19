package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"
	tools "fms-awesome-tools/utils"

	"github.com/spf13/cobra"
)

var (
	start        bool
	dest         string
	lane         string
	auto         bool
	vehicles     int64
	s            int64
	loopNum      int64
	assignedQC   string
	assignedLane string
	noStandby    bool
	onlyStandby  bool
	loadP        float64
	dischargeP   float64
	callP        float64
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
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if assignedQC != "" && !strings.HasPrefix(assignedQC, "PQC") {
			return fmt.Errorf("assigned QC fotmat incorrect")
		}

		if assignedLane != "" {
			v, err := strconv.Atoi(assignedLane)
			if err != nil {
				return err
			}
			if v < 2 || v > 6 || v == 4 {
				return fmt.Errorf("assigned lane must be between 2, 3, 5, 6")
			}
		}
		return nil
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

	if err := service.NewWorkflow(
		loopNum, constants.Activity, vehicles, s, lane, constants.VehicleID, dest, assignedQC, assignedLane, auto, noStandby, onlyStandby, callP,
	).StartWorkflow(); err != nil {
		fmt.Println("failed to start workflow:", err)
		return
	}
}

func init() {
	workflowCmd.Flags().BoolVar(&start, "start", false, "start workflow👻")
	workflowCmd.Flags().Int64VarP(&constants.Activity, "activity", "a", 1, "STANDBY = 1\nMOUNT = 2\nNO_YARD = 5\nOFFLOAD = 6\n")
	workflowCmd.Flags().StringVarP(&constants.VehicleID, "truck", "v", "", "集卡号🚗")
	workflowCmd.Flags().StringVarP(&dest, "destination", "d", "", "任务的目的地; QC: PQC921, Block: Y,V,,TB01,32,32,10, ;🔚")
	workflowCmd.Flags().StringVarP(&lane, "lane", "l", "", "目的地车道")
	workflowCmd.Flags().StringVar(&assignedQC, "assigned-qc", "", "指定作业QC")
	workflowCmd.Flags().StringVar(&assignedLane, "assigned-lane", "", "指定QC的作业车道")
	workflowCmd.Flags().BoolVarP(&auto, "auto-call", "", false, "自动发送call-in🔄️")
	workflowCmd.Flags().BoolVar(&noStandby, "no-standby", false, "禁止Standby任务🔄️")
	workflowCmd.Flags().BoolVar(&onlyStandby, "only-standby", false, "只运行Standby任务🔄️")
	workflowCmd.Flags().Int64VarP(&vehicles, "vehicles", "", 0, "执行workflow的集卡数量 (从APM9001开始编号)")
	workflowCmd.Flags().Float64VarP(&callP, "call-percentage", "c", 100, "自动call-in的概率（1-100）")
	workflowCmd.Flags().Int64Var(&s, "start-num", 0, "执行workflow的集卡起始号")
	workflowCmd.Flags().Int64Var(&loopNum, "loop", 0, "循环执行workflow\n-1: 无限循环\n0: 执行一次\n>0: 执行指定次数\n新任务目的地轮换指定, QC: PQC924-2, 堆场: 随机指定\n")
	workflowCmd.MarkFlagsRequiredTogether("destination", "lane")
	workflowCmd.MarkFlagsRequiredTogether("vehicles", "start-num")
	workflowCmd.MarkFlagsMutuallyExclusive("truck", "vehicles")
	workflowCmd.MarkFlagsMutuallyExclusive("no-standby", "only-standby")
	rootCmd.AddCommand(workflowCmd)
}
