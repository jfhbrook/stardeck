package lib

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
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

func getActiveWindow() (string, error) {
	log.Trace().Msg("kdotool getactivewindow")
	cmd := exec.Command("kdotool", "getactivewindow")
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()

	activeWindow := string(out)

	return activeWindow, err
}

func getWindowName(activeWindow string) (string, error) {
	log.Trace().Msg(fmt.Sprintf("kdotool getwindowname %s", activeWindow))
	cmd := exec.Command("kdotool", "getwindowname", activeWindow)
	cmd.Stderr = os.Stderr

	// TODO: Yields Error: No such object path '/Scripting/Script0'
	out, err := cmd.Output()

	windowName := string(out)

	return windowName, err
}

func (w *window) poll() (*Event, error) {
	activeWindow, err := getActiveWindow()

	if err != nil {
		log.Debug().Err(err).Msg("Error when calling kdotool getactivewindow")
		return nil, err
	}

	windowName, err := getWindowName(activeWindow)

	if err != nil {
		log.Debug().Err(err).Msg("Error when calling kdotool getwindowname")
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
			log.Warn().Msg("Could not poll window")
		} else if event != nil {
			*events <- event
		}

		time.Sleep(duration)
	}
}
