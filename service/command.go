package service

import (
	"slices"

	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"

	// "github.com/jfhbrook/stardeck/loopback"
	"github.com/jfhbrook/stardeck/plusdeck"
)

type commandType int

const (
	setWindowNameCommand    commandType = 0
	setLoopbackCommand                  = 1
	setPlusdeckStateCommand             = 2
)

type command struct {
	Type  commandType
	Value any
}

func newSetWindowNameCommand(name string) *command {
	return &command{
		Type:  setWindowNameCommand,
		Value: name,
	}
}

func newSetLoopbackCommand(managed bool) *command {
	return &command{
		Type:  setLoopbackCommand,
		Value: managed,
	}
}

func newSetPlusdeckStateCommand(state plusdeck.PlusdeckState) *command {
	return &command{
		Type:  setPlusdeckStateCommand,
		Value: state,
	}
}

func plusdeckStateSetter(state *plusdeck.PlusdeckState) func(update plusdeck.PlusdeckState) {
	displayedStates := []plusdeck.PlusdeckState{
		plusdeck.PlayingA,
		plusdeck.PausedA,
		plusdeck.PlayingB,
		plusdeck.PausedB,
		plusdeck.FastForwardingA,
		plusdeck.FastForwardingB,
	}

	return func(update plusdeck.PlusdeckState) {
		if slices.Contains(displayedStates, update) {
			log.Warn().Str("state", update).Msg("TODO: Display plusdeck state")
		} else {
			log.Debug().Str("state", update).Msg("NOTE: Do not display plusdeck state")
		}
		*state = update
	}
}

func CommandRunner(systemConn *dbus.Conn, commands chan *command) {
	windowName := ""
	plusdeckState := plusdeck.Unsubscribed

	sendData := crystalfontzSender(systemConn)

	lb := newLoopbackManager(&plusdeckState)
	setPlusdeckState := plusdeckStateSetter(&plusdeckState)

	for {
		log.Trace().Msg("Waiting for command")
		command := <-commands
		log.Trace().Any("command", command).Msg("Received command")

		switch command.Type {
		case setWindowNameCommand:
			sendData(command.Value.(string))
		case setLoopbackCommand:
			lb.set(command.Value.(bool))
		case setPlusdeckStateCommand:
			setPlusdeckState(command.Value.(plusdeck.PlusdeckState))
		}

		log.Debug().
			Str("windowName", windowName).
			Any("plusdeckState", plusdeckState).
			Msg("State updated")
	}
}
