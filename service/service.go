package service

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"

	"github.com/jfhbrook/stardeck/crystalfontz"
	"github.com/jfhbrook/stardeck/logger"
	"github.com/jfhbrook/stardeck/plusdeck"
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

	lcd := crystalfontz.NewClient(systemConn)

	line1 := newLcdLine(0, "YES THIS IS STARDECK", lcd)
	line2 := newLcdLine(1, "", lcd)

	lb := newLoopbackManager(plusdeck.Unsubscribed)
	pd := newPlusdeckManager(plusdeck.Unsubscribed, line1)
	note := newNotificationManager(line2)

	go serve(commands)
	go listen(systemConn, sessionConn, events)
	go eventHandler(events, commands)
	go signalHandler(lcd)
	CommandRunner(
		line1,
		line2,
		lb,
		pd,
		note,
		commands,
	)
}
