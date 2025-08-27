package lib

import (
	"github.com/godbus/dbus/v5"
)

func AddCrystalfontzMatchSignal(conn *dbus.Conn) error {
	return conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/"),
		dbus.WithMatchInterface("org.jfhbrook.crystalfontz"),
	)
}

func HandleCrystalfontzKeyActivityReport(signal *dbus.Signal, events *chan *Event) {
	// signal.Body[0] should be a byte representation of the report
	// this will need to be unpacked
}
