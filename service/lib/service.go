package lib

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

func Service() error {
	windowPollInterval := 0.2

	sessionConn, err := dbus.ConnectSessionBus()

	if err != nil {
		return err
	}

	systemConn, err := dbus.ConnectSystemBus()

	if err != nil {
		return err
	}

	events := make(chan *Event)

	go ListenToWindow(windowPollInterval, &events)
	go ListenToSignals(systemConn, &events)
	go ListenToNotifications(sessionConn, &events)

	for {
		event := <-events

		switch event.Type {
		case WindowEvent:
			fmt.Println(event)
		case PlusdeckEvent:
			fmt.Println(event)
		case KeyActivityReport:
			fmt.Println(event)
		case Notification:
			fmt.Println(event)
		}
	}
}
