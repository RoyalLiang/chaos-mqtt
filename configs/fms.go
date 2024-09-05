package configs

import (
	"encoding/json"
	"github.com/spf13/viper"
)

var (
	vp    = viper.GetViper()
	Chaos *chaosConfig
)

type chaosConfig struct {
	Product product `json:"product"`
	MQTT    mqtt    `json:"mqtt"`
}

type product struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type mqtt struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func (c *chaosConfig) String() string {
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

func defaultConfig() chaosConfig {
	config := chaosConfig{}
	config.Product = product{
		UUID:    "93a25ee9-0f08-4398-a0e4-72aa28ee1ebf",
		Version: "1.0.0",
		Name:    "chaos",
	}
	config.MQTT = mqtt{
		Address:  "",
		User:     "",
		Password: "",
	}
	return config
}

func init() {
	Chaos = &chaosConfig{}
}
