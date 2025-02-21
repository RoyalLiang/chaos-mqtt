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
)

var APMArrivedCmd = &cobra.Command{
	Use:   "apm_arrived_request",
	Short: "发送 APM arrived message",
	Run: func(cmd *cobra.Command, args []string) {
		if arrivedPosition <= 0 {
			_ = cmd.Help()
		} else {
			location := ""
			switch arrivedPosition {
			case 1:
				location = "Y,V,,"
			case 2:
				location = "P,PQCXXX  _Pre-Ingress"
			case 3:
				location = "P,PQCXXX _Waiting Area"
			case 4:
				location = "P,PQCXXX _Call In Area"
			case 5:
				location = "P,PQCXXX "
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
}
