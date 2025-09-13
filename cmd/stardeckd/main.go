package main

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/config"
	"github.com/jfhbrook/stardeck/logger"
	"github.com/jfhbrook/stardeck/service"
)

var (
	cfgFile   string
	logLevel  string
	logFormat string
	logColor  bool
)

func main() {
	flag.StringVar(&cfgFile, "config-file", "", "Config file (default is /etc/stardeck/stardeck.yml)")
	flag.StringVar(&logLevel, "log-level", "info", "Log level (default is 'info')")
	flag.StringVar(&logFormat, "log-format", logger.JsonFormat, "Log format (default is 'json')")
	flag.BoolVar(&logColor, "log-color", false, "Show logs in color (default is 'false')")

	flag.Parse()

	viper.Set("service.log_level", logLevel)
	viper.Set("service.log_format", logFormat)
	viper.Set("service.log_color", logColor)

	config.InitConfig(cfgFile, config.Service)

	service.Service()
}
