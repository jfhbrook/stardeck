package lib

import (
	// "fmt"
	// "io"
	"os/exec"
)

type WindowWorker struct {
	WindowInterval float64
	ExitTimeout    float64
	ExitInterval   float64
	ExitTryTimes   int
	running        bool
	Window         string
}

func NewWindowWorker() *WindowWorker {
	windowInterval := 0.2
	exitTimeout := 10.0
	exitInterval := 0.1
	exitTryTimes := int(exitTimeout * (1.0 / exitInterval))
	w := WindowWorker{
		WindowInterval: windowInterval,
		ExitTimeout:    exitTimeout,
		ExitInterval:   exitInterval,
		ExitTryTimes:   exitTryTimes,
		running:        true,
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

func (w *WindowWorker) Run(events *chan *Event) {
	for w.running {
		event, err := w.Poll()

		if err != nil {
			panic(err)
		}

		if event != nil {
			*events <- event
		}
	}
}
