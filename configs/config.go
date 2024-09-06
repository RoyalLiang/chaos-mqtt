package configs

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigFile = "chaos.yaml"
)

var (
	ConfigDir string
)

func writeDefaultConfig() error {
	config := defaultConfig()
	vp.Set("product", config.Product)
	vp.Set("mqtt", config.MQTT)
	return vp.WriteConfig()
}

func checkConfig(path string) {
	_, err := os.Stat(path)
	var exist = true
	if os.IsNotExist(err) {
		exist = false
		if err = os.MkdirAll(ConfigDir, 0755); err != nil {
			fmt.Println("配置文件创建失败: ", err)
			os.Exit(1)
		}
		if _, err := os.Create(path); err != nil {
			fmt.Println("配置文件创建失败: ", err)
			os.Exit(1)
		}
	}
	vp.SetConfigFile(path)

	if exist == false {
		if err = writeDefaultConfig(); err != nil {
			fmt.Println("配置写入失败: ", err)
			os.Exit(1)
		}
	}
}

func readConfig() {
	if err := vp.ReadInConfig(); err != nil {
		fmt.Println("配置文件读取失败: ", err)
		os.Exit(1)
	}
	if err := vp.Unmarshal(&Chaos); err != nil {
		fmt.Println("配置文件解析失败: ", err)
		os.Exit(1)
	}
}

func Init() {
	cp := filepath.Join(ConfigDir, ConfigFile)
	checkConfig(cp)
	readConfig()
}

func init() {
	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("获取用户目录失败, ", err)
		os.Exit(1)
	}
	ConfigDir = filepath.Join(path, ".chaos", "config")
}
