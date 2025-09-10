package service

import (
	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"
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
	info := notificationInfo{
		AppName:       payload[0].(string),
		ReplacesId:    payload[1].(uint),
		AppIcon:       payload[2].(string),
		Summary:       payload[3].(string),
		Body:          payload[4].(string),
		Actions:       mapActions(payload[5].([]string)),
		Hints:         payload[6].(map[string]any),
		ExpireTimeout: payload[7].(int32),
	}

	ev := event{Type: notification, Value: info}

	return &ev
}

func eventHandler(events chan *event, commands chan *command) {
	for {
		ev := <-events

		switch ev.Type {
		case windowEvent:
			log.Debug().Any("event", ev).Msg("WindowEvent")
		case plusdeckEvent:
			log.Debug().Any("event", ev).Msg("PlusdeckEvent")
		case keyActivityReport:
			log.Debug().Any("event", ev).Msg("KeyActivityReport")
		case notification:
			log.Debug().Any("event", ev).Msg("Notification")
		}
	}
}

func listen(systemConn *dbus.Conn, sessionConn *dbus.Conn, events chan *event) {
	go listenToSignals(systemConn, events)
	go listenToNotifications(sessionConn, events)
}
