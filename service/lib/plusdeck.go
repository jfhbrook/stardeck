package lib

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"os/exec"
)

type PlusdeckState = string

const (
	PlayingA        PlusdeckState = "PLAYING_A"
	PausedA                       = "PAUSED_A"
	PlayingB                      = "PLAYING_B"
	Subscribed                    = "SUBSCRIBED"
	PausedB                       = "PAUSED_B"
	FastForwardingA               = "FAST_FORWARDING_A"
	FastForwardingB               = "FAST_FORWARDING_B"
	Stopped                       = "STOPPED"
	Ejected                       = "EJECTED"
	Subscribing                   = "SUBSCRIBING"
	Unsubscribing                 = "UNSUBSCRIBING"
	Unsubscribed                  = "UNSUBSCRIBED"
)

func AddPlusdeckMatchSignal(conn *dbus.Conn) error {
	return conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/"),
		dbus.WithMatchInterface("org.jfhbrook.crystalfontz"),
	)
}

func newPlusdeckEvent(state string) *Event {
	e := Event{Type: PlusdeckEvent, Value: state}

	return &e
}

func HandlePlusdeckState(signal *dbus.Signal, events chan *Event) {
	events <- newPlusdeckEvent(signal.Body[0].(string))
}

const (
	DefaultLoopbackSource  string = "alsa_input.pci-0000_00_1f.3.analog-stereo"
	DefaultLoopbackLatency int32  = 1
	DefaultLoopbackVolume  int32  = 10
)

type LoopbackManager struct {
	source  string
	latency int32
	volume  int32
}

func NewLoopbackManager(source string, latency int32, volume int32) *LoopbackManager {
	src := source
	if src == "" {
		src = DefaultLoopbackSource
	}

	lt := latency
	if lt < 0 {
		lt = DefaultLoopbackLatency
	}

	vol := volume
	if vol < 0 {
		vol = DefaultLoopbackVolume
	}

	lb := LoopbackManager{source: src, latency: lt, volume: vol}

	return &lb
}

func (lb *LoopbackManager) Enable() error {
	if err := exec.Command(
		"pactl",
		"load-module",
		"module-loopback",
		fmt.Sprintf("--latency_msec=%d", lb.latency),
	).Run(); err != nil {
		return err
	}

	return exec.Command(
		"pactl",
		"set-source-volume",
		lb.source,
		fmt.Sprintf("%d", lb.volume),
	).Run()
}

func (lb *LoopbackManager) Disable() error {
	return exec.Command(
		"pactl",
		"unload-module",
		"module-loopback",
	).Run()
}
