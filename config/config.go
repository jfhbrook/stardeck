package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/logger"
)

func InitConfig(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/etc/stardeck")
		viper.SetConfigType("yaml")
		viper.SetConfigName("stardeck.yml")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	logger.ConfigureLogger(viper.GetString("log-level"))

	if err != nil {
		log.Debug().Msg(err.Error())
	} else {
		log.Debug().Msg(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
	}
}
