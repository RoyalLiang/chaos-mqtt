package envs

import (
	"github.com/spf13/cobra"
)

var (
	url    string
	module string
	port   string
)

var FMSCmd = &cobra.Command{
	Use:   "fms",
	Short: "FMS模块配置",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	FMSCmd.Flags().StringVarP(&url, "host", "u", "", "模块HOST地址")
	FMSCmd.Flags().StringVarP(&module, "module", "m", "", "模块名称")
	FMSCmd.Flags().StringVarP(&port, "port", "p", "", "模块启动端口")
}
