package topics

import (
	"github.com/spf13/cobra"
)

var (
	manual int64
)

var CAllCmd = &cobra.Command{
	Use:   "call_in_request",
	Short: "发送call_in_request",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	CAllCmd.Flags().Int64VarP(&manual, "manual", "m", 0, "优先级")
}
