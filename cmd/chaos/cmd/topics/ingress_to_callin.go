package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var IngressToCallInCmd = &cobra.Command{
	Use:   "ingress_to_call_in",
	Short: "发送 ingress_to_call_in",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("ingress_to_call_in", constants.IngressToCallIn, generateIngressToCallIn()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateIngressToCallIn() interface{} {
	return constants.IngressToCallInParam{
		VehicleID: constants.VehicleID,
	}
}
