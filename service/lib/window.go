package lib

import (
	"os/exec"
)

type WindowWorker struct {
	WindowInterval float64
	Window         string
}

func NewWindowWorker(interval float64) *WindowWorker {
	w := WindowWorker{
		WindowInterval: interval,
		Window:         "",
	}
	return &w
}

func newWindowEvent(name string) *Event {
	e := Event{Type: WindowEvent, Value: name}

	return &e
}

func (w *WindowWorker) Poll() (*Event, error) {
	activeWindow, err := exec.Command("kdotool", "getactivewindow").Output()

	if err != nil {
		return nil, err
	}

	// TODO: Can I assign to err like that?
	windowName, err := exec.Command("kdotool", "getwindowname", string(activeWindow)).Output()

	if err != nil {
		return nil, err
	}

	next := string(windowName)

	if next != w.Window {
		w.Window = next
		e := newWindowEvent(w.Window)
		return e, nil
	}

	return nil, nil
}

func ListenToWindow(interval float64, events *chan *Event) error {
	w := NewWindowWorker(interval)

	for {
		event, err := w.Poll()

		if err != nil {
			panic(err)
		}

		if event != nil {
			*events <- event
		}
	}
}
