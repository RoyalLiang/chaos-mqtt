package topics

import (
	tools "fms-awesome-tools/utils"
	"fmt"

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
		if err := service.PublishAssignedTopic("route_request", constants.RouteRequest, generateRouteRequestParam()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateRouteRequestParam() interface{} {
	return constants.RouteRequestParam{
		VehicleID:   constants.VehicleID,
		Type:        dType,
		Destination: dest,
	}
}

func init() {
	RouteRequestCmd.Flags().StringVarP(&dType, "type", "t", "", "route 类型")
	RouteRequestCmd.Flags().StringVarP(&dest, "dest", "d", "", "route 目的地")
}
