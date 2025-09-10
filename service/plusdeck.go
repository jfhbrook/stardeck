package service

import (
	"github.com/godbus/dbus/v5"
)

func newPlusdeckEvent(state string) *Event {
	e := Event{Type: PlusdeckEvent, Value: state}

	return &e
}

func HandlePlusdeckState(signal *dbus.Signal, events chan *Event) {
	events <- newPlusdeckEvent(signal.Body[0].(string))
}
