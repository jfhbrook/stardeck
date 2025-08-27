package lib

import (
	"github.com/godbus/dbus/v5"
)

func AddPlusdeckMatchSignal(conn *dbus.Conn) error {
	return conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/"),
		dbus.WithMatchInterface("org.jfhbrook.crystalfontz"),
	)
}

func HandlePlusdeckState(signal *dbus.Signal, events *chan *Event) {
	// signal.Body[0] should be the event name
}
