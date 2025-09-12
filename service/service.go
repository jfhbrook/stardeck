package service

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"

	"github.com/jfhbrook/stardeck/logger"
)

func Service() {
	sessionConn, err := dbus.ConnectSessionBus()

	if err != nil {
		logger.FlagrantError(errors.Wrap(err, "Failed to connect to session bus"))
	}

	defer sessionConn.Close()

	systemConn, err := dbus.ConnectSystemBus()

	if err != nil {
		logger.FlagrantError(errors.Wrap(err, "Failed to connect to system bus"))
	}

	defer systemConn.Close()

	exportIface(sessionConn)

	events := make(chan *event)
	commands := make(chan *command)

	go listen(systemConn, sessionConn, events)
	go eventHandler(events, commands)
	CommandRunner(commands)
}
