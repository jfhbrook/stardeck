package service

import (
	"slices"

	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/plusdeck"
)

type plusdeckManager struct {
	displayedStates []plusdeck.State
	state         plusdeck.State
}

func newPlusdeckManager(state plusdeck.State) *plusdeckManager {
	displayedStates := []plusdeck.State{
		plusdeck.PlayingA,
		plusdeck.PausedA,
		plusdeck.PlayingB,
		plusdeck.PausedB,
		plusdeck.FastForwardingA,
		plusdeck.FastForwardingB,
	}

	return &plusdeckManager{
		displayedStates: displayedStates,
		state: state,
	}
}

func (pd *plusdeckManager) update(state plusdeck.State) {
	if slices.Contains(pd.displayedStates, state) {
		log.Warn().Str("state", state).Msg("TODO: Display plusdeck state")
	} else {
		log.Debug().Str("state", state).Msg("NOTE: Do not display plusdeck state")
	}
	pd.state = state
}
