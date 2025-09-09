package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("stardeck")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/stardeck")
}

func readInConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug().Msg(err.Error())
			return nil
		}
		return err
	}
	return nil
}

func handleConfigFileNotFoundError(err error) bool {
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		log.Debug().Msg(err.Error())
		return true
	}
	return false
}
