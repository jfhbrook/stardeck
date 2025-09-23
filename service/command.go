package service

import (
	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/crystalfontz"
	"github.com/jfhbrook/stardeck/notifications"
	"github.com/jfhbrook/stardeck/plusdeck"
)

type commandType int

const (
	setWindowNameCommand    commandType = 0
	setLoopbackCommand                  = 1
	setPlusdeckStateCommand             = 2
	showNotificationCommand             = 3
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

func newSetLoopbackCommand(managed bool) *command {
	return &command{
		command: setLoopbackCommand,
		Value:   managed,
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

func CommandRunner(systemConn *dbus.Conn, commands chan *command) {
	windowName := ""
	loopbackManaged := viper.GetBool("loopback.managed")
	plusdeckState := plusdeck.Unsubscribed

	lcd := crystalfontz.NewClient(systemConn)

	line1 := newLcdLine(0, "YES THIS IS STARDECK", lcd)
	line2 := newLcdLine(1, "", lcd)

	line1.start()
	line2.start()

	lb := newLoopbackManager(plusdeckState)
	pd := newPlusdeckManager(plusdeckState, line1)
	note := newNotificationManager(line2)

	for {
		log.Trace().Msg("Waiting for command")
		command := <-commands
		log.Trace().Any("command", command).Msg("Received command")

		switch command.command {
		case setWindowNameCommand:
			windowName = command.Value.(string)
		case setLoopbackCommand:
			loopbackManaged = command.Value.(bool)
		case setPlusdeckStateCommand:
			plusdeckState = command.Value.(plusdeck.State)
		case showNotificationCommand:
			note.update(command.Value.(*notifications.NotificationInfo))
		}

		log.Debug().
			Str("windowName", windowName).
			Bool("loopbackManaged", loopbackManaged).
			Any("plusdeckState", plusdeckState).
			Msg("State updated")

		lb.update(loopbackManaged, plusdeckState)
		if !pd.update(plusdeckState) {
			line1.update(windowName)
		}
	}
}
