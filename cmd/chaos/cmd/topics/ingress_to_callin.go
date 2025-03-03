package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"

	"fms-awesome-tools/cmd/chaos/service"
)

var IngressToCallInCmd = &cobra.Command{
	Use:   "ingress_to_call_in",
	Short: "发送 ingress_to_call_in",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("ingress_to_call_in", "", generateIngressToCallIn()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateIngressToCallIn() interface{} {
	return messages.IngressToCallIn{
		APMID: constants.VehicleID,
		Data: messages.IngressToCallInData{
			RouteDag: make([]messages.RouteDag, 0), LaneAvailability: make([]string, 0), RouteType: "G",
		},
	}.String()
}
