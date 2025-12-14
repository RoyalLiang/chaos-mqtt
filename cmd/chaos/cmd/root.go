package cmd

import (
	"fms-awesome-tools/configs"
	tools "fms-awesome-tools/utils"
	"fmt"
	"os"

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
	if configs.Chaos.Logger == nil {
		fmt.Println("logger not initialized, exit...")
		os.Exit(1)
	}
	tools.InitialLogger(configs.Root, *configs.Chaos.Logger)
}
