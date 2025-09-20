package service

import (
	"slices"

	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/plusdeck"
)

func plusdeckStateSetter(state *plusdeck.State) func(update plusdeck.State) {
	displayedStates := []plusdeck.State{
		plusdeck.PlayingA,
		plusdeck.PausedA,
		plusdeck.PlayingB,
		plusdeck.PausedB,
		plusdeck.FastForwardingA,
		plusdeck.FastForwardingB,
	}

	return func(update plusdeck.State) {
		if slices.Contains(displayedStates, update) {
			log.Warn().Str("state", update).Msg("TODO: Display plusdeck state")
		} else {
			log.Debug().Str("state", update).Msg("NOTE: Do not display plusdeck state")
		}
		*state = update
	}
}
