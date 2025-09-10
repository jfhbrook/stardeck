package service

import (
	"github.com/godbus/dbus/v5"
)

func newKeyActivityReportEvent(activity byte) *Event {
	e := Event{Type: KeyActivityReport, Value: activity}

	return &e
}

func HandleKeyActivityReport(signal *dbus.Signal, events chan *Event) {
	events <- newKeyActivityReportEvent(signal.Body[0].(byte))
}
