package cmd

import (
	"fms-awesome-tools/configs"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{

	Use:     "chaos",
	Short:   "***************************************************\nA simple MQTT CLI daemon program for FMS with AVCS.\n***************************************************",
	Long:    "***************************************************\nA simple MQTT CLI daemon program for FMS with AVCS.\n***************************************************",
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
