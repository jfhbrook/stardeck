package plusdeck

import (
	"github.com/pkg/errors"
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
	err := conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/"),
		dbus.WithMatchInterface("org.jfhbrook.plusdeck"),
	)

	if err != nil {
		return errors.Wrap(err, "Failed to match signal for plusdeck state")
	}

	return nil
}
