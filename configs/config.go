package configs

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	ConfigDir  = "configs"
	ConfigFile = ".fms.yaml"
	LogDir     = "logs"
	LogFile    = ".log"
)

var (
	Version   string
	BuildTime string
	Commit    string
)

func ReadConfig(vp *viper.Viper, path string, cfgObj FMSConfig) (*FMSConfig, error) {

	vp.SetConfigFile(path)
	if err := vp.ReadInConfig(); err != nil {
		zap.S().Error("[ Viper Read Error ] ", err)
		return nil, err
	}
	if err := vp.Unmarshal(&cfgObj); err != nil {
		zap.S().Error("[ Viper Parse Error ] ", err)
		return nil, err
	}
	return &cfgObj, nil
}

func WriteConfig(vp *viper.Viper, key string, value interface{}) {
	vp.Set(key, value)
	err := vp.WriteConfig()
	if err != nil {
		zap.S().Warn("[ Config Write Error ] fail to config ", key, ": ", value)
	}
}

func init() {
	viper.SetConfigName(ConfigFile)
	viper.SetConfigType("yaml")
}
