package topics

import (
	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

const locations = "\n1=堆场位置\n2=Pre-Ingress\n3=DWA\n4=CA\n5=QC\n6=REFUEL\n7=PARKING\n8=MAINTENANCE\n9=CALLBACK"

var (
	arrivedPosition int
	aCrane          string
	aBlock          string
)

var APMArrivedCmd = &cobra.Command{
	Use:   "apm_arrived_request",
	Short: "发送 APM arrived message",
	Run: func(cmd *cobra.Command, args []string) {
		if arrivedPosition <= 0 {
			_ = cmd.Help()
		} else {
			if arrivedPosition == 1 && (aBlock == "" || aCrane != "") {
				cobra.CheckErr("目的地与到达类型不匹配")
			}

			craneType := []int{2, 3, 4, 5}
			for _, v := range craneType {
				if v == arrivedPosition && (aBlock != "" || aCrane == "") {
					cobra.CheckErr("目的地与到达类型不匹配")
				}
			}

			location := ""
			switch arrivedPosition {
			case 1:
				location = "Y,V,," + aBlock
			case 2:
				location = fmt.Sprintf("P,%s  _Pre-Ingress", aCrane)
			case 3:
				location = fmt.Sprintf("P,%s  _Waiting Area", aCrane)
			case 4:
				location = fmt.Sprintf("P,%s  _Call In Area", aCrane)
			case 5:
				location = fmt.Sprintf("P,%s   ", aCrane)
			case 6:
				location = "Refuel_X_X"
			case 7:
				location = "MB_XX"
			case 8:
				location = "MB_XX"
			case 9:
				location = "Callback_XX"
			}

			data := generateAPMArrivedRequest(location).String()
			if err := service.PublishAssignedTopic("/apm/apm_arrived_request", "", data); err != nil {
				cobra.CheckErr(err)
			} else {
				fmt.Println(data, " ==> apm_arrived_request")
			}
		}
	},
}

func generateAPMArrivedRequest(location string) messages.APMArrivedRequest {
	return messages.APMArrivedRequest{
		APMID: constants.VehicleID,
		Data: messages.APMArrivedRequestData{
			ID:                 "",
			Location:           location,
			AlternateLaneDock:  "D",
			TargetDockPosition: "3",
			Timestamp:          time.Now().UnixMilli(),
		},
	}
}

func init() {
	APMArrivedCmd.Flags().IntVarP(&arrivedPosition, "position", "p", *new(int), "到达位置"+locations)
	APMArrivedCmd.Flags().StringVar(&aCrane, "crane", "", "到达的QC")
	APMArrivedCmd.Flags().StringVar(&aBlock, "block", "", "到达的block")
	APMArrivedCmd.MarkFlagsMutuallyExclusive("crane", "block")
}
