package service

import (
	"slices"

	"github.com/rs/zerolog/log"
	// "github.com/jfhbrook/stardeck/loopback"
	"github.com/jfhbrook/stardeck/plusdeck"
)

type loopbackManager struct {
	managed       bool
	managedStates []plusdeck.PlusdeckState
	state         *plusdeck.PlusdeckState
}

func newLoopbackManager(state *plusdeck.PlusdeckState) *loopbackManager {
	managedStates := []plusdeck.PlusdeckState{
		plusdeck.PlayingA,
		plusdeck.PausedA,
		plusdeck.PlayingB,
		plusdeck.PausedB,
		plusdeck.FastForwardingA,
		plusdeck.FastForwardingB,
		plusdeck.Stopped,
	}

	return &loopbackManager{
		managed:       false,
		managedStates: managedStates,
		state:         state,
	}
}

func (m *loopbackManager) isManagedState() bool {
	return slices.Contains(m.managedStates, *m.state)
}

func (m *loopbackManager) enable() {
	log.Warn().Msg("TODO: enable loopback")
}

func (m *loopbackManager) disable() {
	log.Warn().Msg("TODO: disable loopback")
}

func (m *loopbackManager) set(managed bool) {
	if managed {
		if m.isManagedState() {
			m.enable()
		} else {
			m.disable()
		}
	}
}
