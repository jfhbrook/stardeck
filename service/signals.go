package service

import (
	"github.com/godbus/dbus/v5"

	"github.com/jfhbrook/stardeck/crystalfontz"
	"github.com/jfhbrook/stardeck/logger"
	"github.com/jfhbrook/stardeck/plusdeck"
)

func ListenToSignals(conn *dbus.Conn, events chan *Event) {
	if err := plusdeck.AddStateMatchSignal(conn); err != nil {
		logger.FlagrantError(err)
	}
	if err := crystalfontz.AddKeyActivityReportMatchSignal(conn); err != nil {
		logger.FlagrantError(err)
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
}
