package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"fms-awesome-tools/configs"
	tools "fms-awesome-tools/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	host    string
	port    string
	cfgFile string
)
var envCmd = &cobra.Command{

	Use:   "env",
	Short: "读取/写入相关配置",
	Long:  `读取/写入相关配置`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("readConfig") {
			readConfig()
		}

		if host != "" && port != "" {
			mqtt := readMQTTConfig()
			mqtt.Address = append(mqtt.Address, host+":"+port)
			configs.WriteConfig(viper.GetViper(), "mqtt", mqtt)
			fmt.Println("已配置的MQTT地址: ")
			for _, v := range mqtt.Address {
				fmt.Println("-", v)
			}
		} else {
			fmt.Println("")
		}
	},
}

func readConfig() {
	fmt.Println("读取配置文件")
}

func readMQTTConfig() *configs.MQTT {
	cp := filepath.Join(tools.GetRootDir(), configs.ConfigDir, configs.ConfigFile)
	config, err := configs.ReadConfig(viper.GetViper(), cp, configs.FMSConfig{})
	if err != nil {
		fmt.Println("配置文件读取失败: ", err)
		os.Exit(1)
	}
	return &config.MQTT
}

func writeConfig(key, value string) {

}

func init() {
	cobra.OnInitialize(initConfig)

	envCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件")
	envCmd.PersistentFlags().StringVarP(&host, "MQTT_HOST", "", "", "设置MQTT Server 地址")
	envCmd.PersistentFlags().StringVarP(&port, "MQTT_PORT", "", "", "设置MQTT Server 端口")

	envCmd.PersistentFlags().BoolP("list", "l", false, "列出所有配置")
	viper.BindPFlag("readConfig", envCmd.PersistentFlags().Lookup("list"))

	rootCmd.AddCommand(envCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home := tools.GetRootDir()

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".fms")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
