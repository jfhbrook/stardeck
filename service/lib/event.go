package lib

type EventType int

const (
	WindowEvent EventType = iota
	PlusdeckEvent
	KeyActivityReport
)

type Event struct {
	Type  EventType
	Value any
}
