package cmd

import (
	"fms-awesome-tools/cmd/chaos/cmd/requests"
	tools "fms-awesome-tools/utils"
	"github.com/spf13/cobra"
)

var (
	server string
)

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "request data to assigned server",
	Long:  tools.CustomTitle("request data to assigned server"),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	requestCmd.PersistentFlags().StringVarP(&server, "server", "s", "", "指定的service")
	requestCmd.MarkFlagsRequiredTogether("server")

	requestCmd.AddCommand(requests.QCPosCmd)

	rootCmd.AddCommand(requestCmd)
}
