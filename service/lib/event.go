package lib

type EventType int

const (
	WindowEvent EventType = iota
	PlusdeckStateEvent
	CrystalfontzKeyActivityReportEvent
)

// TODO: Make this an interface
type Event struct {
	EventType EventType
	Name      string
}
