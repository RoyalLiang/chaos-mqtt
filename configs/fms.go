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
	Product *product      `json:"product"`
	MQTT    *mqtt         `json:"mqtt"`
	FMS     *fms          `json:"fms"`
	Redis   *RedisConfig  `json:"redis"`
	Logger  *LoggerConfig `json:"logger"`
}

type product struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type LoggerConfig struct {
	Name   string `json:"name"`
	Size   int    `json:"size"`
	Backup int    `json:"backup"`
	Level  string `json:"level"`
	Dir    string `json:"dir"`
}

func (l LoggerConfig) String() string {
	s, _ := json.Marshal(l)
	return string(s)
}

type fms struct {
	//Host     string       `json:"host"`
	//Services []FmsService `json:"services"`
	Area         FMSModuleConfig `json:"area"`
	TOS          FMSModuleConfig `json:"tos"`
	Device       FMSModuleConfig `json:"device"`
	CraneManager FMSModuleConfig `json:"crane_manager"`
}

type RedisConfig struct {
	Address  string `json:"address"`
	DB       int    `json:"db"`
	Password string `json:"password"`
}

type FMSModuleConfig struct {
	Address string `json:"address"`
}

type FmsService struct {
	Name    string `json:"name"`
	Address string `json:"address"`
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
	config.Logger = &LoggerConfig{
		Name:   "chaos",
		Size:   500,
		Backup: 10,
		Level:  "info",
		Dir:    "logs",
	}

	config.FMS = &fms{
		Area:         FMSModuleConfig{},
		TOS:          FMSModuleConfig{},
		Device:       FMSModuleConfig{},
		CraneManager: FMSModuleConfig{},
	}
	config.Redis = &RedisConfig{}
	return config
}

func init() {
	Chaos = &chaosConfig{}
	Init()
}
