package envs

import (
	"fms-awesome-tools/configs"
	"fmt"
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
		if url == "" && module == "" && port == "" {
			_ = cmd.Help()
			return
		}

		if url != "" {
			if err := configs.WriteFMSConfig("fms.host", url); err != nil {
				fmt.Println("HOST配置失败: ", err)
			}
		}

		if module != "" && port != "" {
			if err := configs.WriteFMSConfig(fmt.Sprintf("fms.services.%s.%s", module, port), port); err != nil {
				fmt.Println("FMS HOST配置失败:", err)
			}
		}
	},
}

func init() {
	FMSCmd.Flags().StringVarP(&url, "host", "u", "", "模块HOST地址")
	FMSCmd.Flags().StringVarP(&module, "module", "m", "", "模块名称")
	FMSCmd.Flags().StringVarP(&port, "port", "p", "", "模块启动端口")
	FMSCmd.MarkFlagsRequiredTogether("module", "port")
}
