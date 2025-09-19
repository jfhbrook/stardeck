package service

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"

	"github.com/jfhbrook/stardeck/logger"
)

func serve(commands chan *command) {
	conn, err := dbus.ConnectSessionBus()

	if err != nil {
		logger.FlagrantError(errors.Wrap(err, "Failed to connect to session bus"))
	}

	defer conn.Close()

	err = exportIface(conn, commands)

	if err != nil {
		logger.FlagrantError(errors.Wrap(err, "Failed to export interface"))
	}

	select {}
}

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

	events := make(chan *event, 1)
	commands := make(chan *command, 1)

	go serve(commands)
	go listen(systemConn, sessionConn, events)
	go eventHandler(events, commands)
	CommandRunner(systemConn, commands)
}
