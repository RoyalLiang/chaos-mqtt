package cmd

import (
	"fms-awesome-tools/cmd/chaos/cmd/http/area"
	tools "fms-awesome-tools/utils"
	"github.com/spf13/cobra"
)

var (
	ip string
)

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "request data to assigned server",
	Long:  tools.CustomTitle("request data to assigned server"),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	requestCmd.PersistentFlags().StringVarP(&ip, "ip", "", "127.0.0.1", "服务地址")

	requestCmd.AddCommand(area.SetBlockCmd)
	requestCmd.AddCommand(area.ManualModeCmd)
	requestCmd.AddCommand(area.GetVesselCmd)
	rootCmd.AddCommand(requestCmd)
}
