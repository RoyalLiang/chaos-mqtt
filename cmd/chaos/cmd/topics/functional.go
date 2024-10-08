package topics

import (
	"fms-awesome-tools/cmd/chaos/internal/messages"
	"fms-awesome-tools/cmd/chaos/service"
	"fms-awesome-tools/constants"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name string
)

var FunctionalCmd = &cobra.Command{
	Use:   "function_job",
	Short: "发送 refuel/parking/maintenance功能区任务",
	Run: func(cmd *cobra.Command, args []string) {

		topic := ""
		switch name {
		case "refuel", "parking", "maintenance":
			topic = name
		default:
			fmt.Printf("Unknown job type: %s\n", name)
			return
		}

		if err := service.PublishAssignedTopic(topic, "", generateFunctionalJob()); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func generateFunctionalJob() string {
	return messages.FunctionalJob{
		APMID: constants.VehicleID,
		Data:  messages.FunctionalJobData{RouteDag: make([]messages.RouteDag, 0), RouteType: "G"},
	}.String()
}

func init() {
	FunctionalCmd.Flags().StringVarP(&name, "type", "t", "maintenance", "maintenance\nparking\nrefuel")
}
