package lib

import (
	"github.com/godbus/dbus/v5"
)

func Service() {
	interval := 0.2

	configureLogger()

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

	events := make(chan *Event)

	go Listen(systemConn, sessionConn, &events, interval)

	CommandRunner(events)
}
