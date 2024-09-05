package cmd

import (
	"fms-awesome-tools/configs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{

	Use:     "chaos",
	Short:   "A simple CLI daemon program for MQTT with FMS.",
	Long:    `A simple CLI daemon program for MQTT with FMS.`,
	Version: Version(),
	Run:     Run,
}

func Run(cmd *cobra.Command, args []string) {
	if cmd.Name() == "chaos" {
		_ = cmd.Help()
	}
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
