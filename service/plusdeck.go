package service

import (
	"github.com/jfhbrook/stardeck/plusdeck"
)

type plusdeckManager struct {
	states map[plusdeck.State]string
	state  plusdeck.State
	line   *lcdLine
}

func newPlusdeckManager(state plusdeck.State, line *lcdLine) *plusdeckManager {
	states := map[plusdeck.State]string{
		plusdeck.PlayingA:        "Playing Side A",
		plusdeck.PausedA:         "Paused Side A",
		plusdeck.PlayingB:        "Playing Side B",
		plusdeck.PausedB:         "Paused Side B",
		plusdeck.FastForwardingA: "Fast Forwarding Side A (Rewinding Side B)",
		plusdeck.FastForwardingB: "Fast Forwarding Side B (Rewinding Side A)",
	}

	return &plusdeckManager{
		states: states,
		state:  state,
		line:   line,
	}
}

func (pd *plusdeckManager) update(state plusdeck.State) bool {
	pd.state = state
	name, ok := pd.states[state]
	if ok {
		pd.line.update(name)
	}
	return ok
}
