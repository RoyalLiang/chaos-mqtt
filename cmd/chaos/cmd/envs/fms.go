package envs

import (
	"fms-awesome-tools/configs"
	"fms-awesome-tools/constants"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	area   bool
	tos    bool
	device bool
	crane  bool
)

var FMSCmd = &cobra.Command{
	Use:   "fms",
	Short: "FMS模块配置",
	Run: func(cmd *cobra.Command, args []string) {
		if (!area && !tos && !device) || constants.Address == "" {
			_ = cmd.Help()
			return
		}

		cfg := configs.FMSModuleConfig{
			Address: constants.Address,
		}
		if area {
			configs.Chaos.FMS.Area = cfg
		} else if tos {
			configs.Chaos.FMS.TOS = cfg
		} else if device {
			configs.Chaos.FMS.Device = cfg
		} else if crane {
			configs.Chaos.FMS.CraneManager = cfg
		}

		if err := configs.WriteFMSConfig("fms", configs.Chaos.FMS); err != nil {
			cobra.CheckErr(err)
		}
		fmt.Println("配置成功...")
	},
}

func init() {
	FMSCmd.Flags().BoolVar(&area, "area", false, "area-slot 模块")
	FMSCmd.Flags().BoolVar(&tos, "tos", false, "tos-interface 模块")
	FMSCmd.Flags().BoolVar(&device, "device", false, "device-info 模块")
	FMSCmd.Flags().BoolVar(&crane, "crane-manager", false, "crane-manager 模块")
	FMSCmd.MarkFlagsMutuallyExclusive("area", "tos", "device", "crane-manager")
}
