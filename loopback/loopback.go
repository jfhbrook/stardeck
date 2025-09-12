package loopback

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/loopback/parser"
)

const (
	DefaultLoopbackSource  string = "alsa_input.pci-0000_00_1f.3.analog-stereo"
	DefaultLoopbackSink  string = "alsa_output.pci-0000_00_1f.3.analog-stereo"
	DefaultLoopbackLatency int32  = 1
	DefaultLoopbackVolume  int32  = 10
)

type Status struct {
	Source string
	Latency int32
	Volume int32
}

type LoopbackManager struct {
	source  string
	sink string
	latency int32
	volume  int32
}

func NewLoopbackManager(source string, sink string, latency int32, volume int32) *LoopbackManager {
	src := source
	if src == "" {
		src = DefaultLoopbackSource
	}

	snk := sink
	if snk == "" {
		snk = DefaultLoopbackSink
	}

	lt := latency
	if lt < 0 {
		lt = DefaultLoopbackLatency
	}

	vol := volume
	if vol < 0 {
		vol = DefaultLoopbackVolume
	}

	lb := LoopbackManager{source: src, sink: snk, latency: lt, volume: vol}

	return &lb
}

func (lb *LoopbackManager) getModule() (*parser.Module, error) {
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

func (lb *LoopbackManager) getSinkNo() (string, error) {
	output, err := exec.Command(
		"pactl",
		"list",
		"sinks",
		"short",
	).Output()

	if err != nil {
		return "", err
	}

	sinks := strings.Split(string(output), "\n")

	for _, sink := range sinks {
		fields := strings.Fields(sink)
		if len(fields) >= 2 {
			if fields[1] == lb.sink {
			  return fields[0], nil
			}
		}
	}

	return "", errors.New("Sink not found")
}

func (lb *LoopbackManager) getVolume() (int, error) {
	sinkNo, err := lb.getSinkNo()

	if err != nil {
		return -1, err
	}

	output, err := exec.Command(
		"pactl",
		"get-sink-volume",
		sinkNo,
	).Output()

	if err != nil {
		return -1, err
	}

	re := regexp.MustCompile(`\d+`)
	found := re.Find(output)
	return strconv.Atoi(string(found))
}

func (lb *LoopbackManager) Status() (*Status, error) {
	module, err := lb.getModule()

	if err != nil {
		return nil, err
	}

	latencyParam := module.Params["--latency_msec"]
  latency, err := strconv.Atoi(latencyParam)

	if err != nil {
		log.Debug().Msg(err.Error())
		latency = -1
	}

	volume, err := lb.getVolume()

	if err != nil {
		log.Debug().Msg(err.Error())
	}

  st := Status {
		Source: lb.source,
		Latency: int32(latency),
		Volume: int32(volume),
	}

	return &st, nil
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

	return lb.setVolume()
}

func (lb *LoopbackManager) setVolume() error {
	sinkNo, err := lb.getSinkNo()

	if err != nil {
		return err
	}

	return exec.Command(
		"pactl",
		"set-source-volume",
		sinkNo,
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
