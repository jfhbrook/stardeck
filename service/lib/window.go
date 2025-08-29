package lib

import (
	"os/exec"
	"time"
)

type window struct {
	name string
}

func newWindow() *window {
	w := window{name: ""}
	return &w
}

func newWindowEvent(name string) *Event {
	e := Event{Type: WindowEvent, Value: name}

	return &e
}

func (w *window) poll() (*Event, error) {
	activeWindow, err := exec.Command("kdotool", "getactivewindow").Output()

	if err != nil {
		return nil, err
	}

	windowName, err := exec.Command("kdotool", "getwindowname", string(activeWindow)).Output()

	if err != nil {
		return nil, err
	}

	next := string(windowName)

	if next != w.name {
		w.name = next
		e := newWindowEvent(w.name)
		return e, nil
	}

	return nil, nil
}

func ListenToWindow(interval float64, events *chan *Event) error {
	duration := time.Duration(interval * float64(time.Second))
	w := newWindow()

	for {
		event, err := w.poll()

		if err != nil {
			flagrantError(err)
		}

		if event != nil {
			*events <- event
		}

		time.Sleep(duration)
	}
}
