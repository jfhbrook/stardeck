package lib

import (
	"github.com/godbus/dbus/v5"
	"github.com/spf13/viper"
)

func Service() {
	configureLogger()
	initConfig()

	if err := viper.ReadInConfig(); err != nil {
		if !handleConfigFileNotFoundError(err) {
			flagrantError(err)
		}
	}

	interval := viper.GetFloat64("poll_interval")

	sessionConn, err := dbus.ConnectSessionBus()

	if err != nil {
		flagrantError(err)
	}

	defer sessionConn.Close()

	systemConn, err := dbus.ConnectSystemBus()

	if err != nil {
		flagrantError(err)
	}

	defer systemConn.Close()

	exportIface(sessionConn)

	events := make(chan *Event)
	commands := make(chan *Command)

	go Listen(systemConn, sessionConn, events, interval)
	go EventHandler(events, commands)
	CommandRunner(commands)
}
