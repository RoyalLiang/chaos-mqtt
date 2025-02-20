package envs

import (
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	url  string
	name string
	port string
)

var FMSCmd = &cobra.Command{
	Use:   "fms",
	Short: "FMS模块配置",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		if url == "" && name == "" && address == "" {
			_ = cmd.Help()
			return
		}

		if url != "" {
			if err := configs.WriteFMSConfig("fms.host", url); err != nil {
				fmt.Println("HOST配置失败: ", err)
			}
		}

		if name != "" && address != "" {
			cfg := &configs.FmsService{
				Name:    name,
				Address: address,
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
	FMSCmd.Flags().StringVarP(&name, "name", "n", "", "模块名称")
	FMSCmd.Flags().StringVarP(&port, "port", "p", "", "模块启动端口")
	FMSCmd.Flags().StringVarP(&address, "address", "a", "", "模块base地址")
	FMSCmd.MarkFlagsRequiredTogether("name", "address")
}
