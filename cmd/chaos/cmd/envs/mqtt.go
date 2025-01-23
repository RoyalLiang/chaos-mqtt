package envs

import (
	"fmt"

	"fms-awesome-tools/configs"
	tools "fms-awesome-tools/utils"
	"github.com/spf13/cobra"
)

var (
	user    string
	pwd     string
	address string
)

var MQTTCmd = &cobra.Command{
	Use:  "mqtt",
	Long: tools.CustomTitle("写入MQTT相关配置"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mqtt args: ", args)
		if user != "" || pwd != "" || address != "" {
			if err := configs.WriteFMSConfig("mqtt.address", address); err != nil {
				fmt.Println("MQTT地址配置失败:", err)
			}
			if err := configs.WriteFMSConfig("mqtt.user", user); err != nil {
				fmt.Println("用户名配置失败:", err)
			}
			if err := configs.WriteFMSConfig("mqtt.password", pwd); err != nil {
				fmt.Println("密码配置配置失败:", err)
			}

		} else {
			_ = cmd.Help()
		}
	},
}

func init() {
	MQTTCmd.Flags().StringVarP(&user, "user", "u", "", "用户名")
	MQTTCmd.Flags().StringVarP(&pwd, "password", "p", "", "密码")
	MQTTCmd.Flags().StringVarP(&address, "address", "h", "", "服务端地址")
}
