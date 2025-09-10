package plusdeck

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

func AddStateMatchSignal(conn *dbus.Conn) error {
	return conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/"),
		dbus.WithMatchInterface("org.jfhbrook.plusdeck"),
	)
}
