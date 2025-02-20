package area

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	vehicleID string
)

var VehicleCmd = &cobra.Command{
	Use:   "vehicles",
	Short: "获取所有/指定集卡状态",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args: ", args)
		return
	},
}

func init() {
	VehicleCmd.Flags().Bool("keepalive", false, "是否保持刷新(F/5s)")
	VehicleCmd.Flags().StringVarP(&vehicleID, "vehicle", "v", "", "集卡号")
}
