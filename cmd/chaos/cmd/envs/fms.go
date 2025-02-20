package envs

import (
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	url     string
	name    string
	port    string
	baseUrl string
)

var FMSCmd = &cobra.Command{
	Use:   "fms",
	Short: "FMS模块配置",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		if url == "" && name == "" && baseUrl == "" {
			_ = cmd.Help()
			return
		}

		if url != "" {
			if err := configs.WriteFMSConfig("fms.host", url); err != nil {
				fmt.Println("HOST配置失败: ", err)
			}
		}

		if name != "" && baseUrl != "" {
			cfg := &configs.FmsService{
				Name:    name,
				BaseUrl: baseUrl,
			}

			services := configs.Chaos.FMS.Services
			services = append(services, *cfg)
			configs.Chaos.FMS.Services = services
			if err := configs.WriteFMSConfig("fms", configs.Chaos.FMS); err != nil {
				fmt.Println("FMS HOST配置失败:", err)
			}
		}
	},
}

func init() {
	FMSCmd.Flags().StringVarP(&url, "host", "u", "", "FMS HOST地址")
	FMSCmd.Flags().StringVarP(&name, "module", "m", "", "模块名称")
	FMSCmd.Flags().StringVarP(&port, "port", "p", "", "模块启动端口")
	FMSCmd.Flags().StringVarP(&baseUrl, "base-url", "b", "", "模块base地址")
	FMSCmd.MarkFlagsRequiredTogether("module", "moduleUrl")
}
