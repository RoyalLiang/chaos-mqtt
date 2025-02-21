package cmd

import (
	"fms-awesome-tools/cmd/chaos/cmd/http/area"
	tools "fms-awesome-tools/utils"
	"github.com/spf13/cobra"
)

var (
	ip string
)

var fmsCmd = &cobra.Command{
	Use:   "fms",
	Short: "request data from assigned fms server",
	Long:  tools.CustomTitle("request data to assigned server"),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	fmsCmd.PersistentFlags().StringVarP(&ip, "ip", "", "127.0.0.1", "服务地址")

	fmsCmd.AddCommand(area.HatchCoverCmd)
	fmsCmd.AddCommand(area.ManualModeCmd)
	fmsCmd.AddCommand(area.GetVesselCmd)
	fmsCmd.AddCommand(area.VehicleCmd)
	rootCmd.AddCommand(fmsCmd)
}
