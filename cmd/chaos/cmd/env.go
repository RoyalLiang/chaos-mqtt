package cmd

import (
	"fmt"
	"path/filepath"

	"fms-awesome-tools/cmd/chaos/cmd/envs"
	"fms-awesome-tools/configs"
	"fms-awesome-tools/constants"
	tools "fms-awesome-tools/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	list bool
)
var envCmd = &cobra.Command{

	Use:   "env",
	Short: "读取/写入MQTT/Redis/FMS等相关配置",
	Long:  tools.CustomTitle("读取/写入MQTT/Redis/FMS等相关配置"),
	Run: func(cmd *cobra.Command, args []string) {

		if !list {
			_ = cmd.Help()
		} else {
			listConfig()
			return
		}
	},
}

//func parseConfig() (key, value string) {
//	array := strings.Split(config, "=")
//	if len(array) != 2 {
//		fmt.Printf("can not parse config option: <%s>\n", config)
//		os.Exit(1)
//	}
//	return array[0], array[1]
//}

func listConfig() {

	vp := viper.GetViper()
	cp := filepath.Join(configs.ConfigDir, configs.ConfigFile)
	vp.SetConfigFile(cp)
	if err := vp.ReadInConfig(); err != nil {
		fmt.Println("配置读取失败: ", err)
		return
	}

	for _, key := range vp.AllKeys() {
		fmt.Printf("%s=%s\n", key, vp.Get(key))
	}
}

func init() {
	envCmd.AddCommand(envs.MQTTCmd)
	envCmd.AddCommand(envs.FMSCmd)
	envCmd.AddCommand(envs.RedisCmd)

	envCmd.Flags().BoolVarP(&list, "list", "l", false, "列出当前配置列表😎")
	envCmd.PersistentFlags().StringVarP(&constants.Address, "address", "a", "", "服务base url")

	rootCmd.AddCommand(envCmd)
}
