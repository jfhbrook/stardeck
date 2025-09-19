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
	e := event{Type: keyActivityReport, Value: activity}

	return &e
}

func newPlusdeckEvent(state string) *event {
	e := event{Type: plusdeckEvent, Value: state}

	return &e
}

func newNotificationEvent(payload []any) *event {
	info := notifications.NewNotificationInfo(payload)

	ev := event{Type: notification, Value: info}

	return &ev
}

func eventHandler(events chan *event, commands chan *command) {
	windowName := ""
	for {
		ev := <-events

		switch ev.Type {
		case windowEvent:
			if ev.Value.(string) != windowName {
				windowName = ev.Value.(string)
				setWindowNameCmd := makeSetWindowNameCommand(windowName)
				commands <- setWindowNameCmd
			}
		case plusdeckEvent:
			log.Debug().Any("event", ev).Msg("PlusdeckEvent")
		case keyActivityReport:
			log.Debug().Any("event", ev).Msg("KeyActivityReport")
		case notification:
			log.Debug().Any("event", ev).Msg("Notification")
		}
	}
}

func loadInitialState(conn *dbus.Conn, signals chan *dbus.Signal, events chan *event) {
	states := make(chan string)
	go func() {
		pd := plusdeck.NewPlusdeck(conn)
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
		fmt.Printf("%#v\n", message)
		// event := newNotificationEvent(message.Body)

		// *events <- event
	}
}

func listen(systemConn *dbus.Conn, sessionConn *dbus.Conn, events chan *event) {
	go listenToSignals(systemConn, events)
	go listenToNotifications(sessionConn, events)
}
