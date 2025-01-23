package envs

import (
	tools "fms-awesome-tools/utils"
	"github.com/spf13/cobra"
)

var FMSCmd = &cobra.Command{
	Use:  "fms",
	Long: tools.CustomTitle("FMS模块配置"),
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {

}
