package area

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	ingress int64
	egress  int64
)

var ManualModeCmd = &cobra.Command{
	Use:   "manual_mode",
	Short: "手动模式",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("args: ", args)
		} else {
			_ = cmd.Help()
		}
	},
}

func init() {
	ManualModeCmd.Flags().Int64VarP(&ingress, "ingress", "i", 0, "指定的ingress wharf mark")
	ManualModeCmd.Flags().Int64VarP(&egress, "egress", "e", 0, "指定的egress wharf mark")
}
