package service

import (
	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"

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

func newSetPlusdeckStateCommand(state plusdeck.State) *command {
	return &command{
		Type:  setPlusdeckStateCommand,
		Value: state,
	}
}

func CommandRunner(systemConn *dbus.Conn, commands chan *command) {
	windowName := ""
	loopbackManaged := false
	plusdeckState := plusdeck.Unsubscribed

	sendData := crystalfontzSender(systemConn)

	lb := newLoopbackManager(plusdeckState)
	// setPlusdeckState := plusdeckStateSetter(&plusdeckState)

	for {
		log.Trace().Msg("Waiting for command")
		command := <-commands
		log.Trace().Any("command", command).Msg("Received command")

		switch command.Type {
		case setWindowNameCommand:
			sendData(command.Value.(string))
		case setLoopbackCommand:
			loopbackManaged := command.Value.(bool)
			lb.update(loopbackManaged, plusdeckState)
		case setPlusdeckStateCommand:
			plusdeckState := command.Value.(plusdeck.State)
			// setPlusdeckState(plusdeckState)
			lb.update(loopbackManaged, plusdeckState)
		}

		log.Debug().
			Str("windowName", windowName).
			Bool("loopbackManaged", loopbackManaged).
			Any("plusdeckState", plusdeckState).
			Msg("State updated")
	}
}
