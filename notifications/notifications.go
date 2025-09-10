package notifications

import (
	"github.com/godbus/dbus/v5"
)

type NotificationInfo struct {
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

func NewNotificationInfo(payload []any) *NotificationInfo {
	info := NotificationInfo{
		AppName:       payload[0].(string),
		ReplacesId:    payload[1].(uint),
		AppIcon:       payload[2].(string),
		Summary:       payload[3].(string),
		Body:          payload[4].(string),
		Actions:       mapActions(payload[5].([]string)),
		Hints:         payload[6].(map[string]any),
		ExpireTimeout: payload[7].(int32),
	}

	return &info
}

func Eavesdrop(conn *dbus.Conn) (chan *dbus.Message, error) {
	rules := []string{
		"type='method_call',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
	}
	var flag uint = 0

	call := conn.BusObject().Call("org.freedesktop.DBus.Monitoring.BecomeMonitor", 0, rules, flag)

	if call.Err != nil {
		return nil, call.Err
	}

	messages := make(chan *dbus.Message)

	conn.Eavesdrop(messages)

	return messages, nil
}
