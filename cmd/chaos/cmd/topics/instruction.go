package topics

import (
	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	taskID   string
	opAction string
)

var InstructionCmd = &cobra.Command{
	Use:   "instruction",
	Short: "发送 mount_instruction/offload_instruction",
	Run: func(cmd *cobra.Command, args []string) {
		if taskID == "" && opAction == "" {
			_ = cmd.Help()
			return
		}

		if opAction != "mount" && opAction != "offload" {
			cobra.CheckErr("action not in valid list")
		}

		var data interface{}
		switch opAction {
		case "mount":
			data = generateMountInstruction().String()
		case "offload":
			data = generateOffloadInstruction().String()
		default:
			_ = cmd.Help()
		}

		if err := service.PublishAssignedTopic(fmt.Sprintf("%s_instruction", opAction), "", data); err != nil {
			cobra.CheckErr(err)
		} else {
			fmt.Println(data, " ==> ", fmt.Sprintf("%s_instruction", opAction))
		}
	},
}

func generateOffloadInstruction() messages.OffloadInstruction {
	return messages.OffloadInstruction{
		APMID: constants.VehicleID,
		Data: messages.OffloadInstructionData{
			ID: taskID, Timestamp: time.Now().UnixMilli(), CntrNumber: "FFFF 0000000", Message: "",
		},
	}

}

func generateMountInstruction() messages.MountInstruction {
	return messages.MountInstruction{
		APMID: constants.VehicleID,
		Data: messages.MountInstructionData{
			ID: taskID, Timestamp: time.Now().UnixMilli(), OperationalType: "IY", JobType: "M", OperationalGroup: "IYO50OO01",
			CntrNumber: "FFFF 0000000", CntrSize: "", CntrWeight: 20000, CntrCategory: "GP", CntrStatus: "F", DGGroup: "",
			IMOClass: "", ReeferTemperature: "", CntrLocationOnAPM: 3, SourceLocation: "", DestLocation: "",
			OffloadSequence: "00", Cone: 0, OperationalQCSequence: 0, OperationalJobSequence: 0, TrailerPosition: "",
			WeightClass: "U", Urgent: "N", DG: "N", LiftType: 1, Message: "",
		},
	}
}

func init() {
	InstructionCmd.Flags().StringVarP(&taskID, "id", "i", "", "任务ID")
	InstructionCmd.Flags().StringVarP(&opAction, "action", "a", "", "动作, 可选:\nmount/offload\n")
	InstructionCmd.MarkFlagsRequiredTogether("id", "action")
}
