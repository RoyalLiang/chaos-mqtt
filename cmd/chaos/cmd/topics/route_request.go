package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/internal/messages"
	tools "fms-awesome-tools/utils"

	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"
)

var (
	dType string
	dest  string
)

var RouteRequestCmd = &cobra.Command{
	Use:   "route_request",
	Short: "发送 route_request",
	Long:  tools.CustomTitle("发送 route_request"),
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PublishAssignedTopic("route_request", "", generateRouteRequestParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateRouteRequestParam() string {
	return messages.RouteRequest{
		ApmId: constants.VehicleID,
		Data: messages.RouteRequestData{
			Type: dType,
			Data: fmt.Sprintf("{\"timestamp\":1725350535153,\"id\":\"MAAPM839103092024160215\",\"map_version\":\"PPT-456-20220620-20240112\",\"dest_location\":\"%s\"}", dest),
		},
	}.String()
}

func init() {
	RouteRequestCmd.Flags().StringVarP(&dType, "type", "t", "", "route 类型")
	RouteRequestCmd.Flags().StringVarP(&dest, "dest", "d", "", "route 目的地")
}
