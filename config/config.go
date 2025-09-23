package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/logger"
)

const (
	Cli     int = 0
	Service     = 1
)

func configureLogger(appType int) {
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
}

func watchConfig(appType int) {
	viper.OnConfigChange(func(e fsnotify.Event) {
		configureLogger(appType)
	})

	viper.WatchConfig()
}

func InitConfig(cfgFile string, appType int) {
	viper.SetDefault("cli.log_level", "info")
	viper.SetDefault("cli.log_format", logger.PrettyFormat)
	viper.SetDefault("cli.log_color", true)

	viper.SetDefault("service.log_format", logger.JsonFormat)
	viper.SetDefault("service.log_level", "info")
	viper.SetDefault("service.log_color", true)

	viper.SetDefault("crystalfontz.width", 16)
	viper.SetDefault("crystalfontz.pause", 5.0)
	viper.SetDefault("crystalfontz.tick", 0.3)

	viper.SetDefault("loopback.managed", true)
	viper.SetDefault("loopback.source", "alsa_input.pci-0000_00_1f.3.analog-stereo")
	viper.SetDefault("loopback.latency", 1)
	viper.SetDefault("loopback.volume", 10)

	viper.SetDefault("notifications.timeout", 15.0)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/etc/stardeck")
		viper.SetConfigType("yaml")
		viper.SetConfigName("stardeck.yml")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	configureLogger(appType)

	if err != nil {
		log.Debug().Msg(err.Error())
	} else {
		log.Debug().Msg(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
		watchConfig(appType)
	}
}
