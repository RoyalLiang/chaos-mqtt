package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fms-awesome-tools/cmd/chaos/cmd/envs"
	"fms-awesome-tools/configs"
	tools "fms-awesome-tools/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configOptions = []string{"user", "password", "address"}
	config        string
	list          bool
)
var envCmd = &cobra.Command{

	Use:   "env",
	Short: "读取/写入相关配置",
	Long:  tools.CustomTitle("读取/写入相关配置"),
	//Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if list {
			listConfig()
			return
		}
		//if !list {
		//	_ = cmd.Help()
		//} else {
		//	listConfig()
		//	return
		//}

		//if config != "" {
		//	key, value := parseConfig()
		//	if allow := func() bool {
		//		for _, v := range configOptions {
		//			if v == key {
		//				return true
		//			}
		//		}
		//		return false
		//	}(); allow == false {
		//		fmt.Println("only support options blew: ")
		//		for _, v := range configOptions {
		//			fmt.Println(v)
		//		}
		//		return
		//	}
		//	if err := configs.WriteFMSConfig("mqtt."+key, value); err != nil {
		//		fmt.Println("write config error:", err)
		//	}
		//}
	},
}

func parseConfig() (key, value string) {
	array := strings.Split(config, "=")
	if len(array) != 2 {
		fmt.Printf("can not parse config option: <%s>\n", config)
		os.Exit(1)
	}
	return array[0], array[1]
}

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

	envCmd.Flags().BoolVarP(&list, "list", "l", false, "列出当前配置列表")

	rootCmd.AddCommand(envCmd)
}
