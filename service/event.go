package service

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/crystalfontz"
	"github.com/jfhbrook/stardeck/logger"
	"github.com/jfhbrook/stardeck/notifications"
	"github.com/jfhbrook/stardeck/plusdeck"
)

type eventType int

const (
	windowEvent eventType = iota
	plusdeckEvent
	keyActivityReport
	notification
)

type event struct {
	Type  eventType
	Value any
}

func newKeyActivityReportEvent(activity byte) *event {
	return &event{Type: keyActivityReport, Value: activity}
}

func newPlusdeckEvent(state string) *event {
	return &event{Type: plusdeckEvent, Value: state}
}

func newNotificationEvent(payload []any) *event {
	info := notifications.NewNotificationInfo(payload)

	return &event{Type: notification, Value: info}
}

func eventHandler(events chan *event, commands chan *command) {
	windowName := ""
	plusdeckState := plusdeck.Unsubscribed

	for {
		log.Trace().Msg("Waiting for event")
		ev := <-events
		log.Trace().Any("event", ev).Msg("Received event")

		switch ev.Type {
		case windowEvent:
			if ev.Value.(string) != windowName {
				windowName = ev.Value.(string)
				commands <- newSetWindowNameCommand(windowName)
			}
		case plusdeckEvent:
			if ev.Value.(plusdeck.State) != plusdeckState {
				plusdeckState = ev.Value.(plusdeck.State)
				commands <- newSetPlusdeckStateCommand(plusdeckState)
			}
		case keyActivityReport:
			log.Debug().Any("event", ev).Msg("keyActivityReport")
		case notification:
			notification := ev.Value.(*notifications.NotificationInfo)
			commands <- newDisplayNotificationCommand(notification)
		}
	}
}

func loadInitialState(conn *dbus.Conn, signals chan *dbus.Signal, events chan *event) {
	states := make(chan string)
	go func() {
		pd := plusdeck.NewClient(conn)
		state, err := pd.CurrentState()

		if err != nil {
			log.Debug().Err(err).Msg("Error while pulling current state")
			return
		}

		states <- state
	}()

	for {
		select {
		case state := <-states:
			events <- newPlusdeckEvent(state)
			return
		case signal := <-signals:
			events <- mapSignal(signal)
			if signal.Name == "org.jfhbrook.plusdeck.State" {
				return
			}
		}
	}
}

func listenToSignals(conn *dbus.Conn, events chan *event) {
	if err := plusdeck.AddStateMatchSignal(conn); err != nil {
		logger.FlagrantError(err)
	}
	if err := crystalfontz.AddKeyActivityReportMatchSignal(conn); err != nil {
		logger.FlagrantError(err)
	}

	signals := make(chan *dbus.Signal, 1)
	conn.Signal(signals)

	loadInitialState(conn, signals, events)

	for signal := range signals {
		events <- mapSignal(signal)
	}
}

func mapSignal(signal *dbus.Signal) *event {
	switch signal.Name {
	case "org.jfhbrook.plusdeck.State":
		return newPlusdeckEvent(signal.Body[0].(string))
	case "org.jfhbrook.crystalfontz.KeyActivityReports":
		return newKeyActivityReportEvent(signal.Body[0].(byte))
	}
	panic(fmt.Sprintf("Unknown signal %s", signal.Name))
}

func listenToNotifications(conn *dbus.Conn, events chan *event) {
	messages, err := notifications.Eavesdrop(conn)

	if err != nil {
		logger.FlagrantError(err)
	}

	for message := range messages {
		if message.Type == dbus.TypeMethodCall {
			event := newNotificationEvent(message.Body)

			events <- event
		}
	}
}

func listen(systemConn *dbus.Conn, sessionConn *dbus.Conn, events chan *event) {
	go listenToSignals(systemConn, events)
	go listenToNotifications(sessionConn, events)
}
