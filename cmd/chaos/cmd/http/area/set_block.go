package area

import (
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	op    string
	start int64
	end   int64
)

var SetBlockCmd = &cobra.Command{
	Use:   "op_wharf_block",
	Short: "添加/删除QC后大梁block",
	Run: func(cmd *cobra.Command, args []string) {
		if op == "" && start <= 0 && end <= 0 {
			_ = cmd.Help()
		} else {
			config := configs.Chaos.FMS
			if config.Host == "" {
				fmt.Println("configuration FMS service HOST first...")
				return
			}

		}
	},
}

func init() {
	SetBlockCmd.Flags().StringVarP(&op, "operate", "o", "", "操作选项, 可选: add/delete")
	SetBlockCmd.Flags().Int64VarP(&start, "start", "s", 0, "开始的wharf mark")
	SetBlockCmd.Flags().Int64VarP(&end, "end", "e", 0, "结束的wharf mark")
	SetBlockCmd.MarkFlagsRequiredTogether("operate", "start", "end")
}
