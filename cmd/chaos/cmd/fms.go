package cmd

import (
	fms "fms-awesome-tools/cmd/chaos/cmd/fms"
	"fms-awesome-tools/cmd/chaos/cmd/http/area"
	tools "fms-awesome-tools/utils"

	"github.com/spf13/cobra"
)

var fmsCmd = &cobra.Command{
	Use:   "fms",
	Short: "与FMS模块进行交互, 接收数据/配置数据",
	Long:  tools.CustomTitle("与FMS模块进行交互, 接收数据/配置数据"),
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	//fmsCmd.PersistentFlags().StringVarP(&ip, "ip", "", "127.0.0.1", "服务地址")

	fmsCmd.AddCommand(area.HatchCoverCmd)
	fmsCmd.AddCommand(area.ManualModeCmd)
	fmsCmd.AddCommand(area.GetVesselCmd)
	fmsCmd.AddCommand(fms.VehicleCmd)
	fmsCmd.AddCommand(fms.CraneMoveCmd)
	fmsCmd.AddCommand(fms.OperateCmd)
	rootCmd.AddCommand(fmsCmd)
}
