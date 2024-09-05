package cmd

import (
	"fms-awesome-tools/configs"
	tools "fms-awesome-tools/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{

	Use:     "chaos",
	Short:   tools.CustomTitle("A simple MQTT CLI daemon program for FMS with AVCS."),
	Long:    tools.CustomTitle("A simple MQTT CLI daemon program for FMS with AVCS."),
	Version: Version(),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Name() == "chaos" {
			_ = cmd.Help()
		}
	},
}

func Version() string {
	return "1.0.0"
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	configs.Init()
}
