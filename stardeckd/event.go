package main

import (
	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"
)

type EventType int

const (
	WindowEvent EventType = iota
	PlusdeckEvent
	KeyActivityReport
	Notification
)

type Event struct {
	Type  EventType
	Value any
}

func EventHandler(events chan *Event, commands chan *Command) {
	for {
		event := <-events

		switch event.Type {
		case WindowEvent:
			log.Debug().Any("event", event).Msg("WindowEvent")
		case PlusdeckEvent:
			log.Debug().Any("event", event).Msg("PlusdeckEvent")
		case KeyActivityReport:
			log.Debug().Any("event", event).Msg("KeyActivityReport")
		case Notification:
			log.Debug().Any("event", event).Msg("Notification")
		}
	}
}

func Listen(systemConn *dbus.Conn, sessionConn *dbus.Conn, events chan *Event) {
	go ListenToSignals(systemConn, events)
	go ListenToNotifications(sessionConn, events)
}
