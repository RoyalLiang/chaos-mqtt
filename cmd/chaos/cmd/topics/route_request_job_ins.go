package topics

import (
	"fmt"

	"fms-awesome-tools/cmd/chaos/internal/messages"

	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"

	"fms-awesome-tools/cmd/chaos/service"
)

const activities = "STANDBY = 1\nMOUNT = 2\nNO_YARD = 5\nOFFLOAD = 6\nOFFLOAD = 7\nOFFLOAD = 8\n"

var (
	destination   string
	container     int64
	lane          string
	targetDockPos string
	liftSize      int64
	quantity      int64
	apmDirection  string
	dBlock        string
	dSlot         string
)

var RouteJobCmd = &cobra.Command{
	Use:   "route_request_job_instruction",
	Short: "发送 route_request_job_instruction",
	Run: func(cmd *cobra.Command, args []string) {
		if targetDockPos != "1" && targetDockPos != "3" && targetDockPos != "5" {
			fmt.Printf("未知的作业位置: %s\n", targetDockPos)
			return
		}

		switch constants.Activity {
		case 1, 5:
			break
		case 2, 3, 4, 6, 7, 8:
			if container >= 40 && quantity > 1 {
				cobra.CheckErr(fmt.Sprintf("箱尺寸 %d 与箱数量 %d 不匹配\n", container, quantity))
			}
		default:
			cobra.CheckErr(fmt.Sprintf("未知的任务类型: %d\n", constants.Activity))
		}

		d := destination
		if dBlock != "" {
			d = fmt.Sprintf("Y,V,,%s,%s,%s,10,   ", dBlock, dSlot, dSlot)
		}

		if err := service.PublishAssignedTopic("route_request_job_instruction", "", messages.GenerateRouteRequestJob(d, lane, apmDirection, targetDockPos, liftSize, container, quantity)); err != nil {
			fmt.Println("error to publish: ", err)
		} else {
			fmt.Println("success to publish")
		}
	},
}

func init() {
	RouteJobCmd.Flags().Int64VarP(&constants.Activity, "activity", "a", 1, activities)
	RouteJobCmd.Flags().StringVarP(&destination, "destination", "d", "", "QC任务目的地, eg: PQC921")
	RouteJobCmd.Flags().Int64VarP(&container, "container-size", "c", 40, "箱尺寸")
	RouteJobCmd.Flags().StringVarP(&lane, "lane", "l", "2", "任务目的地车道")
	RouteJobCmd.Flags().StringVarP(&targetDockPos, "target-dock-position", "t", "5", "任务目的点位; 1: 前箱, 3: 后箱, 5: 双20/单40")
	RouteJobCmd.Flags().StringVar(&apmDirection, "vehicle-direction", "S", "集卡方向; S: 船尾进, B: 船头进")
	RouteJobCmd.Flags().StringVarP(&dBlock, "block", "b", "", "堆场号, eg: TB01")
	RouteJobCmd.Flags().StringVarP(&dSlot, "slot", "x", "", "堆场贝位, eg: 32")
	RouteJobCmd.Flags().Int64VarP(&liftSize, "lift-size", "s", 1, "吊具尺寸; 1: 单20, 2: 双20, 3: 单40/45")
	RouteJobCmd.Flags().Int64VarP(&quantity, "container-quantity", "n", 1, "集装箱数量")
	RouteJobCmd.MarkFlagsMutuallyExclusive("destination", "block")
	RouteJobCmd.MarkFlagsRequiredTogether("block", "slot")
}
