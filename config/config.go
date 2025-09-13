package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/logger"
)

const (
	Cli     int = 0
	Service     = 1
)

func InitConfig(cfgFile string, appType int) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/etc/stardeck")
		viper.SetConfigType("yaml")
		viper.SetConfigName("stardeck.yml")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	level := "info"
	format := logger.JsonFormat
	color := false

	switch appType {
	case Cli:
		level = viper.GetString("cli.log_level")
		format = viper.GetString("cli.log_format")
		color = viper.GetBool("cli.log_color")
	case Service:
		level = viper.GetString("service.log_level")
		format = viper.GetString("service.log_format")
		color = viper.GetBool("service.log_color")
	}

	logger.ConfigureLogger(level, format, color)

	if err != nil {
		log.Debug().Msg(err.Error())
	} else {
		log.Debug().Msg(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
	}
}
