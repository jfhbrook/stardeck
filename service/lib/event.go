package lib

type EventType int

const (
	WindowEvent EventType = iota
	PlusdeckEvent
	KeyActivityReport
	Notification
)

type Event struct {
	Type  EventType
	Value any
}
