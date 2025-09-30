package service

import (
	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/notifications"
	"github.com/jfhbrook/stardeck/plusdeck"
)

type commandType int

const (
	setWindowNameCommand    commandType = 0
	setPlusdeckStateCommand             = 1
	showNotificationCommand             = 2
)

type command struct {
	command commandType
	Value   any
}

func newSetWindowNameCommand(name string) *command {
	return &command{
		command: setWindowNameCommand,
		Value:   name,
	}
}

func newSetPlusdeckStateCommand(state plusdeck.State) *command {
	return &command{
		command: setPlusdeckStateCommand,
		Value:   state,
	}
}

func newDisplayNotificationCommand(notification *notifications.NotificationInfo) *command {
	return &command{
		command: showNotificationCommand,
		Value:   notification,
	}
}

func CommandRunner(
	line1 *lcdLine,
	line2 *lcdLine,
	lb *loopbackManager,
	pd *plusdeckManager,
	note *notificationManager,
	commands chan *command,
) {
	windowName := ""
	plusdeckState := plusdeck.Unsubscribed

	for {
		log.Debug().Msg("Waiting for command")
		command := <-commands
		log.Debug().Any("command", command).Msg("Received command")

		switch command.command {
		case setWindowNameCommand:
			windowName = command.Value.(string)
		case setPlusdeckStateCommand:
			plusdeckState = command.Value.(plusdeck.State)
		case showNotificationCommand:
			note.update(command.Value.(*notifications.NotificationInfo))
		}

		log.Debug().
			Str("windowName", windowName).
			Bool("loopbackManaged", lb.managed).
			Any("plusdeckState", plusdeckState).
			Msg("State updated")

		lb.update(plusdeckState)
		if !pd.update(plusdeckState) {
			line1.update(windowName)
		}
	}
}
