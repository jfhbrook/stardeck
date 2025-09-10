package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("stardeck")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/stardeck")
}

func HandleConfigFileNotFoundError(err error) bool {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		log.Debug().Msg(err.Error())
		return true
	}
	return false
}
