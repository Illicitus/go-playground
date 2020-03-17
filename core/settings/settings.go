package settings

import (
	"github.com/spf13/viper"
	"log"
)

var settings *viper.Viper

func Init(env string) {
	settings = viper.New()
	settings.AddConfigPath("core/settings/")
	settings.SetConfigType("json")
	settings.SetConfigName(env)
	err := settings.ReadInConfig()
	if err != nil {
		log.Panic("Can't read config file")
	}
}

func GetSettings() *viper.Viper {
	return settings
}
