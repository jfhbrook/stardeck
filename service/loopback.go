package service

import (
	"slices"

	"github.com/jfhbrook/stardeck/loopback"
	"github.com/jfhbrook/stardeck/plusdeck"
	"github.com/rs/zerolog/log"
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
		managed:       false,
		enabledStates: enabledStates,
		state:         state,
		manager:       loopback.NewManager(),
		ch:            make(chan loopbackSettings),
	}

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

	m.managed = managed
	m.state = state
}

func (m *loopbackManager) worker() {
	for settings := range m.ch {
		m.toggle(settings.managed, settings.state)
	}
}

func (m *loopbackManager) update(managed bool, state plusdeck.State) {
	m.ch <- loopbackSettings{managed, state}
}
