package envs

import (
	"fmt"

	"fms-awesome-tools/configs"
	"github.com/spf13/cobra"
)

var (
	user    string
	pwd     string
	address string
)

var MQTTCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "写入MQTT相关配置",
	Run: func(cmd *cobra.Command, args []string) {
		if address == "" && pwd == "" && user == "" {
			_ = cmd.Help()
			return
		}

		if address != "" {
			if err := configs.WriteFMSConfig("mqtt.address", address); err != nil {
				fmt.Println("MQTT地址配置失败:", err)
			}
		}

		if user != "" {
			if err := configs.WriteFMSConfig("mqtt.user", user); err != nil {
				fmt.Println("用户名配置失败:", err)
			}
		}

		if pwd != "" {
			if err := configs.WriteFMSConfig("mqtt.password", pwd); err != nil {
				fmt.Println("密码配置配置失败:", err)
			}
		}
	},
}

func init() {
	MQTTCmd.Flags().StringVarP(&user, "user", "u", "", "用户名")
	MQTTCmd.Flags().StringVarP(&pwd, "password", "p", "", "密码")
	MQTTCmd.Flags().StringVarP(&address, "address", "a", "", "服务端地址(host:port)")
}
