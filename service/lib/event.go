package lib

type EventType int

const (
	WindowEvent EventType = iota
)

type Event struct {
	EventType EventType
	Name      string
}
