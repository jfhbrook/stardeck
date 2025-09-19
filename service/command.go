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

func CommandRunner(systemConn *dbus.Conn, commands chan *command) {
	unmanagedStates := []plusdeck.PlusdeckState{
		plusdeck.Subscribed,
		plusdeck.Stopped,
		plusdeck.Ejected,
		plusdeck.Subscribing,
		plusdeck.Unsubscribing,
		plusdeck.Unsubscribed,
	}

	windowName := ""
	loopbackManaged := false
	plusdeckState := plusdeck.Unsubscribed

	// loopbackManager := loopback.NewLoopbackManager("", -1, -1)
	lcd := crystalfontz.NewCrystalfontz(systemConn)

	for {
		log.Trace().Msg("Waiting for command")
		command := <-commands
		log.Debug().Any("command", command).Msg("Received command")

		switch command.Type {
		case setWindowNameCommand:
			updated := command.Value.(string)
			if updated != windowName {
				lcd.SendData(0, 0, []byte(updated[0:16]), -1.0, -1)
			}
			windowName = updated
		case setLoopbackCommand:
			loopbackManaged = command.Value.(bool)
		case setPlusdeckStateCommand:
			plusdeckState = command.Value.(plusdeck.PlusdeckState)
		}

		if loopbackManaged {
			if slices.Contains(unmanagedStates, plusdeckState) {
				log.Debug().Msg("TODO: disable loopback")
			} else {
				log.Debug().Msg("TODO: enable loopback")
			}
		}

		log.Debug().
			Str("windowName", windowName).
			Bool("loopbackManaged", loopbackManaged).
			Any("plusdeckState", plusdeckState).
			Msg("State updated")
	}
}
