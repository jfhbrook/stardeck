package service

import (
	"slices"

	"github.com/jfhbrook/stardeck/plusdeck"
)

type plusdeckManager struct {
	displayedStates []plusdeck.State
	state           plusdeck.State
	line            *lcdLine
}

func newPlusdeckManager(state plusdeck.State, line *lcdLine) *plusdeckManager {
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
		state:           state,
		line:            line,
	}
}

func (pd *plusdeckManager) isDisplaying() bool {
	return slices.Contains(pd.displayedStates, pd.state)
}

func (pd *plusdeckManager) update(state plusdeck.State) bool {
	pd.state = state
	displaying := pd.isDisplaying()
	if displaying {
		pd.line.update(state)
	}
	return displaying
}
