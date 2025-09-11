package loopback

import (
	"fmt"
	"os/exec"

	"github.com/jfhbrook/stardeck/loopback/parser"
)

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

func (lb *LoopbackManager) Status() (*parser.Module, error) {
	output, err := exec.Command(
		"pactl",
		"list",
		"modules",
		"short",
	).Output()

	if err != nil {
		return nil, err
	}

	return parser.ParseModuleOutput(output)
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
