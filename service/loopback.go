package service

import (
	"slices"

	"github.com/rs/zerolog/log"
	// "github.com/jfhbrook/stardeck/loopback"
	"github.com/jfhbrook/stardeck/plusdeck"
)

func loopbackManager(state *plusdeck.PlusdeckState, managed *bool) func(update bool) {
	// loopbackManager := loopback.NewLoopbackManager("", -1, -1)
	managedStates := []plusdeck.PlusdeckState{
		plusdeck.PlayingA,
		plusdeck.PausedA,
		plusdeck.PlayingB,
		plusdeck.PausedB,
		plusdeck.FastForwardingA,
		plusdeck.FastForwardingB,
		plusdeck.Stopped,
	}

	return func(update bool) {
		if update {
			if slices.Contains(managedStates, *state) {
				log.Warn().Msg("TODO: enable loopback")
			} else {
				log.Warn().Msg("TODO: disable loopback")
			}
		}
		*managed = update
	}
}
