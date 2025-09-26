package service

import (
	"slices"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/loopback"
	"github.com/jfhbrook/stardeck/plusdeck"
)

type loopbackSettings struct {
	managed bool
	state   plusdeck.State
}

type loopbackManager struct {
	managed       bool
	enabledStates []plusdeck.State
	state         plusdeck.State
	manager       *loopback.Manager
	ch            chan loopbackSettings
}

func newLoopbackManager(state plusdeck.State) *loopbackManager {
	enabledStates := []plusdeck.State{
		plusdeck.PlayingA,
		plusdeck.PausedA,
		plusdeck.PlayingB,
		plusdeck.PausedB,
		plusdeck.FastForwardingA,
		plusdeck.FastForwardingB,
		plusdeck.Stopped,
	}

	m := &loopbackManager{
		managed:       viper.GetBool("loopback.managed"),
		enabledStates: enabledStates,
		state:         state,
		manager:       loopback.NewManager(),
		ch:            make(chan loopbackSettings),
	}

	m.onConfigChange()
	go m.worker()

	return m
}

func (m *loopbackManager) enable() {
}

func (m *loopbackManager) disable() {
	m.manager.Disable()
	log.Warn().Msg("TODO: disable loopback")
}

func (m *loopbackManager) toggle(managed bool, state plusdeck.State) {
	if managed && slices.Contains(m.enabledStates, state) {
		log.Debug().Msg("Enabling loopback")
		m.manager.Enable()
	} else {
		log.Debug().Msg("Disabling loopback")
		m.manager.Disable()
	}
}

func (m *loopbackManager) worker() {
	for settings := range m.ch {
		m.toggle(settings.managed, settings.state)
	}
}

func (m *loopbackManager) update(state plusdeck.State) {
	m.state = state
	m.ch <- loopbackSettings{m.managed, state}
}

func (m *loopbackManager) onConfigChange() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		managed := viper.GetBool("loopback.managed")

		if managed != m.managed {
			log.Trace().Bool("managed", managed).Msg("Registered new loopback config")

			m.managed = managed
			m.ch <- loopbackSettings{managed, m.state}
		}
	})
}
