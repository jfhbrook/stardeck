package service

import (
	"slices"

	"github.com/rs/zerolog/log"
	// "github.com/jfhbrook/stardeck/loopback"
	"github.com/jfhbrook/stardeck/plusdeck"
)

type loopbackManager struct {
	managed       bool
	enabledStates []plusdeck.PlusdeckState
	state         plusdeck.PlusdeckState
	managedCh     chan bool
	stateCh       chan plusdeck.PlusdeckState
}

func newLoopbackManager(state plusdeck.PlusdeckState) *loopbackManager {
	enabledStates := []plusdeck.PlusdeckState{
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
		managedCh:     make(chan bool),
		stateCh:       make(chan plusdeck.PlusdeckState),
	}

	go m.worker()

	return m
}

func (m *loopbackManager) enable() {
	log.Warn().Msg("TODO: enable loopback")
}

func (m *loopbackManager) disable() {
	log.Warn().Msg("TODO: disable loopback")
}

func (m *loopbackManager) toggle(managed bool, state plusdeck.PlusdeckState) {
	if managed && slices.Contains(m.enabledStates, state) {
		m.enable()
	} else {
		m.disable()
	}

	m.managed = managed
	m.state = state
}

func (m *loopbackManager) worker() {
	for {
		select {
		case managed := <-m.managedCh:
			m.toggle(managed, m.state)
		case state := <-m.stateCh:
			m.toggle(m.managed, state)
		}
	}
}

func (m *loopbackManager) update(managed bool, state plusdeck.PlusdeckState) {
	if managed != m.managed {
		m.managedCh <- managed
	}

	if state != m.state {
		m.stateCh <- state
	}
}
