package envs

import (
	"fmt"

	"fms-awesome-tools/constants"

	"github.com/spf13/cobra"

	"fms-awesome-tools/configs"
)

var (
	user string
	pwd  string
)

var MQTTCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "写入MQTT相关配置",
	Run: func(cmd *cobra.Command, args []string) {
		if constants.Address == "" && pwd == "" && user == "" {
			_ = cmd.Help()
			return
		}

		if constants.Address != "" {
			if err := configs.WriteFMSConfig("mqtt.address", constants.Address); err != nil {
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
}
