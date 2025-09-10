package service

import (
	"fmt"
	"github.com/godbus/dbus/v5"

	"github.com/jfhbrook/stardeck/logger"
)

type notificationInfo struct {
	AppName       string
	ReplacesId    uint
	AppIcon       string
	Summary       string
	Body          string
	Actions       map[string]string
	Hints         map[string]any
	ExpireTimeout int32
}

func mapActions(raw []string) map[string]string {
	var actions map[string]string
	i := 0

	for i < len(raw) {
		actions[raw[i]] = raw[i+1]
		i += 2
	}

	return actions
}

func listenToNotifications(conn *dbus.Conn, events chan *event) {
	rules := []string{
		"type='method_call',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
	}
	var flag uint = 0

	call := conn.BusObject().Call("org.freedesktop.DBus.Monitoring.BecomeMonitor", 0, rules, flag)

	if call.Err != nil {
		logger.FlagrantError(call.Err)
	}

	messages := make(chan *dbus.Message)

	conn.Eavesdrop(messages)

	for message := range messages {
		fmt.Printf("%#v\n", message)
		// event := newNotificationEvent(message.Body)

		// *events <- event
	}
}
