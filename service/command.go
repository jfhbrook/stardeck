package service

import (
	"slices"

	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/crystalfontz"
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
	cmd := command{
		Type:  setWindowNameCommand,
		Value: name,
	}

	return &cmd
}

func newSetLoopbackCommand(managed bool) *command {
	cmd := command{
		Type:  setLoopbackCommand,
		Value: managed,
	}

	return &cmd
}

func newSetPlusdeckStateCommand(state plusdeck.PlusdeckState) *command {
	cmd := command{
		Type:  setPlusdeckStateCommand,
		Value: state,
	}

	return &cmd
}

func stringToData(text string) []byte {
	data := []byte(text)
	if len(data) > 16 {
		data = data[0:16]
	}

	return data
}

func windowNameSetter(lcd *crystalfontz.Crystalfontz, name *string) func(update string) {
	return func(update string) {
		if update != *name {
			lcd.SendData(0, 0, stringToData(update), -1.0, -1)
			*name = update
		}
	}
}

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
	loopbackManaged := false
	plusdeckState := plusdeck.Unsubscribed

	lcd := crystalfontz.NewCrystalfontz(systemConn)

	setWindowName := windowNameSetter(lcd, &windowName)
	manageLoopback := loopbackManager(&plusdeckState, &loopbackManaged)
	setPlusdeckState := plusdeckStateSetter(&plusdeckState)

	for {
		log.Trace().Msg("Waiting for command")
		command := <-commands
		log.Trace().Any("command", command).Msg("Received command")

		switch command.Type {
		case setWindowNameCommand:
			setWindowName(command.Value.(string))
		case setLoopbackCommand:
			manageLoopback(command.Value.(bool))
		case setPlusdeckStateCommand:
			setPlusdeckState(command.Value.(plusdeck.PlusdeckState))
		}

		log.Debug().
			Str("windowName", windowName).
			Bool("loopbackManaged", loopbackManaged).
			Any("plusdeckState", plusdeckState).
			Msg("State updated")
	}
}
