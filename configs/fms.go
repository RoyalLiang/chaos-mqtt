package configs

import (
	"encoding/json"
	"fmt"

	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	vp        = viper.GetViper()
	FMSConfig *fmsConfig
)

type fmsConfig struct {
	Product Product `json:"product"`
	MQTT    MQTT    `json:"mqtt"`
}

type Product struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type MQTT struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func (c *fmsConfig) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

func WriteFMSConfig(key string, value interface{}) error {
	vp.Set(key, value)
	if err := vp.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func init() {

	FMSConfig = &fmsConfig{}
	cp := filepath.Join(ConfigDir, ConfigFile)
	_, err := os.Stat(cp)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(ConfigDir, 0755); err != nil {
			fmt.Println("配置文件创建失败: ", err)
			os.Exit(1)
		}
		if _, err := os.Create(cp); err != nil {
			fmt.Println("配置文件创建失败: ", err)
			os.Exit(1)
		}
	}
	vp.SetConfigFile(cp)
	//vp.SetConfigFile("C:\\Users\\westwell\\projects\\go\\fms-awesome-tools\\configs\\.fms.yaml")

	if err := vp.ReadInConfig(); err != nil {
		fmt.Println("配置文件读取失败: ", err)
		os.Exit(1)
	}
	if err := vp.Unmarshal(&FMSConfig); err != nil {
		fmt.Println("配置文件解析失败: ", err)
		os.Exit(1)
	}
}
