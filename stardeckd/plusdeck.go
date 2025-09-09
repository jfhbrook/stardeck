package main

import (
	"github.com/godbus/dbus/v5"
)

type PlusdeckState = string

const (
	PlayingA        PlusdeckState = "PLAYING_A"
	PausedA                       = "PAUSED_A"
	PlayingB                      = "PLAYING_B"
	Subscribed                    = "SUBSCRIBED"
	PausedB                       = "PAUSED_B"
	FastForwardingA               = "FAST_FORWARDING_A"
	FastForwardingB               = "FAST_FORWARDING_B"
	Stopped                       = "STOPPED"
	Ejected                       = "EJECTED"
	Subscribing                   = "SUBSCRIBING"
	Unsubscribing                 = "UNSUBSCRIBING"
	Unsubscribed                  = "UNSUBSCRIBED"
)

func AddPlusdeckMatchSignal(conn *dbus.Conn) error {
	return conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/"),
		dbus.WithMatchInterface("org.jfhbrook.plusdeck"),
	)
}

func newPlusdeckEvent(state string) *Event {
	e := Event{Type: PlusdeckEvent, Value: state}

	return &e
}

func HandlePlusdeckState(signal *dbus.Signal, events chan *Event) {
	events <- newPlusdeckEvent(signal.Body[0].(string))
}
