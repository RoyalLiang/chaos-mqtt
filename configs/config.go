package configs

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	ConfigDir  = "C:\\Program Files\\fms-tool\\configs"
	ConfigFile = "fms-tool.yaml"
	LogDir     = "logs"
	LogFile    = ".log"
)

var (
	Version   string
	BuildTime string
	Commit    string
)

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
