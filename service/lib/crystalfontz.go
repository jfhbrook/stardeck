package lib

import (
	"github.com/godbus/dbus/v5"
)

type KeyActivity = byte

const (
	KeyUpPress      KeyActivity = 1
	KeyDownPress                = 2
	KeyLeftPress                = 3
	KeyRightPress               = 4
	KeyEnterPress               = 5
	KeyExitPress                = 6
	KeyUpRelease                = 7
	KeyDownRelease              = 8
	KeyLeftRelease              = 9
	KeyRightRelease             = 10
	KeyEnterRelease             = 11
	KeyExitRelease              = 12
)

func AddCrystalfontzMatchSignal(conn *dbus.Conn) error {
	return conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/"),
		dbus.WithMatchInterface("org.jfhbrook.crystalfontz.KeyActivityReports"),
	)
}

func newKeyActivityReportEvent(activity byte) *Event {
	e := Event{Type: KeyActivityReport, Value: activity}

	return &e
}

func HandleKeyActivityReport(signal *dbus.Signal, events *chan *Event) {
	*events <- newKeyActivityReportEvent(signal.Body[0].(byte))
}
