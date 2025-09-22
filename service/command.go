package service

import (
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/crystalfontz"
	"github.com/jfhbrook/stardeck/notifications"
	"github.com/jfhbrook/stardeck/plusdeck"
)

type commandType int

const (
	setWindowNameCommand              commandType = 0
	setLoopbackCommand                            = 1
	setPlusdeckStateCommand                       = 2
	displayNotificationCommand                    = 3
	stopDisplayingNotificationCommand             = 4
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

func newDisplayNotificationCommand(notification *notifications.NotificationInfo) *command {
	return &command{
		Type:  displayNotificationCommand,
		Value: notification,
	}
}

func newStopDisplayingNotificationCommand() *command {
	return &command{
		Type:  stopDisplayingNotificationCommand,
		Value: nil,
	}
}

func CommandRunner(systemConn *dbus.Conn, commands chan *command) {
	windowName := ""
	loopbackManaged := false
	plusdeckState := plusdeck.Unsubscribed

	lcd := crystalfontz.NewClient(systemConn)

	line1 := newLcdLine(0, "Hello!", lcd)
	line2 := newLcdLine(1, "", lcd)

	line1.start()
	line2.start()

	lb := newLoopbackManager(plusdeckState)
	pd := newPlusdeckManager(plusdeckState, line1)

	notificationTimeout := time.Duration(
		viper.GetFloat64("notifications.timeout") * float64(time.Second),
	)

	displayNotification := func(info *notifications.NotificationInfo) {
		text := info.Summary

		if len(info.Body) > 0 {
			text += (" - " + info.Body)
		}

		log.Debug().Str("text", text).Msg("Displaying notification")

		line2.update(text)

		go func() {
			time.Sleep(notificationTimeout)
			log.Debug().Msg("Stop displaying notification")
			commands <- newStopDisplayingNotificationCommand()
		}()
	}

	for {
		log.Trace().Msg("Waiting for command")
		command := <-commands
		log.Trace().Any("command", command).Msg("Received command")

		switch command.Type {
		case setWindowNameCommand:
			windowName = command.Value.(string)
		case setLoopbackCommand:
			loopbackManaged = command.Value.(bool)
		case setPlusdeckStateCommand:
			plusdeckState = command.Value.(plusdeck.State)
		case displayNotificationCommand:
			displayNotification(command.Value.(*notifications.NotificationInfo))
		case stopDisplayingNotificationCommand:
			line2.update("")
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
