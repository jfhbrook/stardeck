package service

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"

	"github.com/jfhbrook/stardeck/logger"
)

func serve() {
	conn, err := dbus.ConnectSessionBus()

	if err != nil {
		logger.FlagrantError(errors.Wrap(err, "Failed to connect to session bus"))
	}

	defer conn.Close()

	err = exportIface(conn)

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

	go serve()

	events := make(chan *event)
	commands := make(chan *command)

	go listen(systemConn, sessionConn, events)
	go eventHandler(events, commands)
	CommandRunner(commands)
}
