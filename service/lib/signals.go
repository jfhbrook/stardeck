package lib

import (
	"github.com/godbus/dbus/v5"
)

func ListenToSignals(conn *dbus.Conn, events *chan *Event) error {
	if err := AddPlusdeckMatchSignal(conn); err != nil {
		return err
	}
	if err := AddCrystalfontzMatchSignal(conn); err != nil {
		return err
	}

	signals := make(chan *dbus.Signal, 10)
	conn.Signal(signals)

	for signal := range signals {
		switch signal.Name {
		case "org.jfhbrook.plusdeck.State":
			HandlePlusdeckState(signal, events)
		case "org.jfhbrook.crystalfontz.KeyActivityReports":
			HandleKeyActivityReport(signal, events)
		}
	}

	return nil
}
