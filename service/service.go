package service

import (
	"github.com/godbus/dbus/v5"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/config"
	"github.com/jfhbrook/stardeck/logger"
)

func Service() {
	logger.ConfigureLogger()
	config.InitConfig()

	if err := viper.ReadInConfig(); err != nil {
		if !config.HandleConfigFileNotFoundError(err) {
			logger.FlagrantError(err)
		}
	}

	sessionConn, err := dbus.ConnectSessionBus()

	if err != nil {
		logger.FlagrantError(err)
	}

	defer sessionConn.Close()

	systemConn, err := dbus.ConnectSystemBus()

	if err != nil {
		logger.FlagrantError(err)
	}

	defer systemConn.Close()

	exportIface(sessionConn)

	events := make(chan *Event)
	commands := make(chan *Command)

	go Listen(systemConn, sessionConn, events)
	go EventHandler(events, commands)
	CommandRunner(commands)
}
