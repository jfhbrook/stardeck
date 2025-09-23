package loopback

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	pkgerrors "github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/loopback/module"
)

const (
	DefaultLoopbackSource  string = "alsa_input.pci-0000_00_1f.3.analog-stereo"
	DefaultLoopbackLatency int32  = 1
	DefaultLoopbackVolume  int32  = 10
)

type Status struct {
	Enabled bool
	Source  string
	Latency int32
	Volume  int32
}

type Manager struct {
	source  string
	latency int32
	volume  int32
}

func NewManager() *Manager {
	src := viper.GetString("loopback.source")
	lt := viper.GetInt32("loopback.latency")
	vol := viper.GetInt32("loopback.volume")

	lb := Manager{source: src, latency: lt, volume: vol}

	return &lb
}

func (lb *Manager) getModule() (*module.Module, error) {
	output, err := exec.Command(
		"pactl",
		"list",
		"modules",
		"short",
	).Output()

	if err != nil {
		return nil, err
	}

	return module.Parse(output)
}

func (lb *Manager) parseSources(output string) (string, error) {
	sources := strings.Split(output, "\n")

	for _, source := range sources {
		fields := strings.Fields(source)
		if len(fields) >= 2 {
			if fields[1] == lb.source {
				return fields[0], nil
			}
		}
	}

	return "", pkgerrors.Wrap(errors.New("Sink not found"), "Sink not found")
}

func (lb *Manager) getSourceNo() (string, error) {
	output, err := exec.Command(
		"pactl",
		"list",
		"sources",
		"short",
	).Output()

	if err != nil {
		return "", pkgerrors.Wrap(err, "Failed to list sources")
	}

	return lb.parseSources(string(output))
}

func parseVolume(output []byte) (int, error) {
	re := regexp.MustCompile(`\d+`)
	found := re.Find(output)
	return strconv.Atoi(string(found))
}

func (lb *Manager) getVolume() (int, error) {
	sourceNo, err := lb.getSourceNo()

	if err != nil {
		return -1, pkgerrors.Wrap(err, "Failed to get source volume")
	}

	output, err := exec.Command(
		"pactl",
		"get-source-volume",
		sourceNo,
	).Output()

	if err != nil {
		return -1, pkgerrors.Wrap(err, "Failed to get source volume")
	}

	volume, err := parseVolume(output)

	if err != nil {
		return -1, pkgerrors.Wrap(err, "Failed to get source volume")
	}

	return volume, nil
}

func (lb *Manager) IsEnabled() (bool, error) {
	mod, err := lb.getModule()

	if err != nil {
		return false, pkgerrors.Wrap(err, "Failed to get status")
	}

	return mod != nil, nil

}

func (lb *Manager) Status() (*Status, error) {
	mod, err := lb.getModule()

	if err != nil {
		return nil, pkgerrors.Wrap(err, "Failed to get status")
	}

	if mod == nil {
		st := Status{
			Enabled: false,
			Source:  lb.source,
			Latency: int32(-1),
			Volume:  int32(-1),
		}
		return &st, nil
	}

	latencyParam := mod.Params["--latency_msec"]
	latency, err := strconv.Atoi(latencyParam)

	if err != nil {
		log.Debug().Msg(err.Error())
		latency = -1
	}

	volume, err := lb.getVolume()

	if err != nil {
		log.Debug().Msg(err.Error())
	}

	st := Status{
		Enabled: true,
		Source:  lb.source,
		Latency: int32(latency),
		Volume:  int32(volume),
	}

	return &st, nil
}

func (lb *Manager) Enable() error {
	isEnabled, err := lb.IsEnabled()

	if err != nil {
		return err
	}

	if isEnabled {
		return nil
	}

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

func (lb *Manager) setVolume() error {
	sourceNo, err := lb.getSourceNo()

	if err != nil {
		return err
	}

	return exec.Command(
		"pactl",
		"set-source-volume",
		sourceNo,
		fmt.Sprintf("%d", lb.volume),
	).Run()
}

func (lb *Manager) Disable() error {
	isEnabled, err := lb.IsEnabled()

	if err != nil {
		return err
	}

	if !isEnabled {
		return nil
	}

	return exec.Command(
		"pactl",
		"unload-module",
		"module-loopback",
	).Run()
}
