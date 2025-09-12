package main

import (
	flag "github.com/spf13/pflag"

	"github.com/jfhbrook/stardeck/config"
	"github.com/jfhbrook/stardeck/logger"
	"github.com/jfhbrook/stardeck/service"
)

var (
	cfgFile  string
	logLevel string
)

func main() {
	flag.StringVar(&cfgFile, "config-file", "", "Config file (default is /etc/stardeck/stardeck.yml)")
	flag.StringVar(&logLevel, "log-level", "info", "Log level (default is 'info')")

	flag.Parse()

	logger.ConfigureLogger(logLevel)
	config.InitConfig(cfgFile)

	service.Service()
}
