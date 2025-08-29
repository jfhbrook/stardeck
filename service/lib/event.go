package lib

import (
	"github.com/godbus/dbus/v5"
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

func Listen(systemConn *dbus.Conn, sessionConn *dbus.Conn, events *chan *Event, interval float64) {
	go ListenToWindow(interval, events)
	go ListenToSignals(systemConn, events)
	go ListenToNotifications(sessionConn, events)
}
