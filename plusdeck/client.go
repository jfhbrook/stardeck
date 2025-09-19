package plusdeck

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"
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

type Plusdeck struct {
	object dbus.BusObject
}

func NewPlusdeck(conn *dbus.Conn) *Plusdeck {
	obj := conn.Object("org.jfhbrook.plusdeck", "/")
	lcd := Plusdeck{object: obj}
	return &lcd
}

func (p *Plusdeck) CurrentState() (string, error) {
	prop, err := p.object.GetProperty("org.jfhbrook.plusdeck.CurrentState")

	if err != nil {
		return "", err
	}

	return prop.Value().(string), nil
}
