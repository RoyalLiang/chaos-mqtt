package envs

import (
	"fms-awesome-tools/configs"
	"fms-awesome-tools/constants"
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
		if url == "" && name == "" && constants.Address == "" {
			_ = cmd.Help()
			return
		}

		if url != "" {
			if err := configs.WriteFMSConfig("fms.host", url); err != nil {
				fmt.Println("HOST配置失败: ", err)
			}
		}

		if name != "" && constants.Address != "" {
			cfg := &configs.FmsService{
				Name:    name,
				Address: constants.Address,
			}

			services := configs.Chaos.FMS.Services
			found := false
			for i, service := range services {
				if service.Name == name {
					services[i] = *cfg
					found = true
					break
				}
			}

			if !found {
				services = append(services, *cfg)
			}

			configs.Chaos.FMS.Services = services
			if err := configs.WriteFMSConfig("fms", configs.Chaos.FMS); err != nil {
				fmt.Println("FMS HOST配置失败:", err.Error())
			}
		}

		fmt.Println("配置成功...")
	},
}

func init() {
	FMSCmd.Flags().StringVarP(&url, "host", "u", "", "FMS HOST地址")
	FMSCmd.Flags().StringVarP(&name, "name", "n", "", "模块名称")
	FMSCmd.Flags().StringVarP(&port, "port", "p", "", "模块启动端口")
}
