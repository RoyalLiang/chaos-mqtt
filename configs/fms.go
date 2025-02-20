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
	Product *product `json:"product"`
	MQTT    *mqtt    `json:"mqtt"`
	FMS     *fms     `json:"fms"`
}

type product struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type fms struct {
	Host     string       `json:"host"`
	Services []FmsService `json:"services"`
}

type FmsService struct {
	Name    string `json:"name"`
	BaseUrl string `json:"baseUrl"`
	//Port string `json:"port"`
}

func (f fms) String() string {
	v, _ := json.Marshal(f)
	return string(v)
}

func (s FmsService) String() string {
	v, _ := json.Marshal(s)
	return string(v)
}

type mqtt struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func (m mqtt) String() string {
	v, _ := json.Marshal(m)
	return string(v)
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
	config := chaosConfig{
		FMS: &fms{},
	}
	config.Product = &product{
		UUID:    "93a25ee9-0f08-4398-a0e4-72aa28ee1ebf",
		Version: "1.0.0",
		Name:    "chaos",
	}
	config.MQTT = &mqtt{
		Address:  "",
		User:     "",
		Password: "",
	}

	config.FMS.Services = []FmsService{}
	config.FMS.Services = append(config.FMS.Services, FmsService{
		Name:    "area",
		BaseUrl: "http://127.0.0.1:8888",
	})
	return config
}

func init() {
	Chaos = &chaosConfig{}
	Init()
}
