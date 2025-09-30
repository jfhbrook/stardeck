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
	defaultLogLevel := "info"
	defaultLogColor := true
	defaultPause := 3.0

	viper.SetDefault("cli.log_level", defaultLogLevel)
	viper.BindEnv("cli.log_level", "STARDECK_CLI_LOG_LEVEL")
	viper.SetDefault("cli.log_format", logger.PrettyFormat)
	viper.BindEnv("cli.log_format", "STARDECK_CLI_LOG_FORMAT")
	viper.SetDefault("cli.log_color", defaultLogColor)
	viper.BindEnv("cli.log_color", "STARDECK_CLI_LOG_COLOR")

	viper.SetDefault("service.log_level", defaultLogLevel)
	viper.BindEnv("service.log_level", "STARDECK_SERVICE_LOG_LEVEL")
	viper.SetDefault("service.log_format", logger.JsonFormat)
	viper.BindEnv("service.log_format", "STARDECK_SERVICE_LOG_FORMAT")
	viper.SetDefault("service.log_color", defaultLogColor)
	viper.BindEnv("service.log_color", "STARDECK_SERVICE_LOG_COLOR")

	viper.SetDefault("crystalfontz.width", 16)
	viper.BindEnv("crystalfontz.width", "STARDECK_CRYSTALFONTZ_WIDTH")
	viper.SetDefault("crystalfontz.pause", defaultPause)
	viper.BindEnv("crystalfontz.pause", "STARDECK_CRYSTALFONTZ_PAUSE")
	viper.SetDefault("crystalfontz.tick", 0.3)
	viper.BindEnv("crystalfontz.tick", "STARDECK_CRYSTALFONTZ_TICK")

	viper.SetDefault("loopback.managed", true)
	viper.BindEnv("loopback.managed", "STARDECK_LOOPBACK_MANAGED")
	viper.SetDefault("loopback.source", "alsa_input.pci-0000_00_1f.3.analog-stereo")
	viper.BindEnv("loopback.source", "STARDECK_LOOPBACK_SOURCE")
	viper.SetDefault("loopback.latency", 1)
	viper.BindEnv("loopback.latency", "STARDECK_LOOPBACK_LATENCY")
	viper.SetDefault("loopback.volume", 10)
	viper.BindEnv("loopback.volume", "STARDECK_LOOPBACK_VOLUME")

	viper.SetDefault("notifications.timeout", 15.0)
	viper.BindEnv("notifications.timeout", "STARDECK_NOTIFICATIONS_TIMEOUT")
	viper.SetDefault("notifications.min_wait", defaultPause)
	viper.BindEnv("notifications.min_wait", "STARDECK_NOTIFICATIONS_MIN_WAIT")
	viper.SetDefault("notifications.max_queue_length", 5)
	viper.BindEnv("notifications.max_queue_length", "STARDECK_NOTIFICATIONS_MAX_QUEUE_LENGTH")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/etc/stardeck")
		viper.SetConfigType("yaml")
		viper.SetConfigName("stardeck.yml")
	}

	viper.SetEnvPrefix("stardeck")
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
