package notifications

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"
)

type NotificationInfo struct {
	AppName       string
	ReplacesId    uint32
	AppIcon       string
	Summary       string
	Body          string
	Actions       map[string]string
	Hints         map[string]dbus.Variant
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
		ReplacesId:    payload[1].(uint32),
		AppIcon:       payload[2].(string),
		Summary:       payload[3].(string),
		Body:          payload[4].(string),
		Actions:       mapActions(payload[5].([]string)),
		Hints:         payload[6].(map[string]dbus.Variant),
		ExpireTimeout: payload[7].(int32),
	}

	return &info
}

// Initial "bad" message:
/*
&dbus.Message{Type:0x4, Flags:0x1, Headers:map[dbus.HeaderField]dbus.Variant{0x1:dbus.Variant{sig:dbus.Signature{str:"o"}, value:"/org/freedesktop/DBus"}, 0x2:dbus.Variant{sig:dbus.Signature{str:"s"}, value:"org.freedesktop.DBus"}, 0x3:dbus.Variant{sig:dbus.Signature{str:"s"}, value:"NameLost"}, 0x6:dbus.Variant{sig:dbus.Signature{str:"s"}, value:":1.197"}, 0x7:dbus.Variant{sig:dbus.Signature{str:"s"}, value:"org.freedesktop.DBus"}, 0x8:dbus.Variant{sig:dbus.Signature{str:"g"}, value:dbus.Signature{str:"s"}}}, Body:[]interface {}{":1.197"}, serial:0xffffffff}
*/
// Actual notification:
/*
&dbus.Message{Type:0x1, Flags:0x0, Headers:map[dbus.HeaderField]dbus.Variant{0x1:dbus.Variant{sig:dbus.Signature{str:"o"}, value:"/org/freedesktop/Notifications"}, 0x2:dbus.Variant{sig:dbus.Signature{str:"s"}, value:"org.freedesktop.Notifications"}, 0x3:dbus.Variant{sig:dbus.Signature{str:"s"}, value:"Notify"}, 0x6:dbus.Variant{sig:dbus.Signature{str:"s"}, value:":1.184"}, 0x7:dbus.Variant{sig:dbus.Signature{str:"s"}, value:":1.201"}, 0x8:dbus.Variant{sig:dbus.Signature{str:"g"}, value:dbus.Signature{str:"susssasa{sv}i"}}}, Body:[]interface {}{"notify-send", 0x0, "", "hello", "there", []string{}, map[string]dbus.Variant{"sender-pid":dbus.Variant{sig:dbus.Signature{str:"x"}, value:270289}, "urgency":dbus.Variant{sig:dbus.Signature{str:"y"}, value:0x1}}, -1}, serial:0x9}
*/
func Eavesdrop(conn *dbus.Conn) (chan *dbus.Message, error) {
	rules := []string{
		"type='method_call',member='Notify',path='/org/freedesktop/Notifications',interface='org.freedesktop.Notifications'",
	}
	var flag uint = 0

	call := conn.BusObject().Call("org.freedesktop.DBus.Monitoring.BecomeMonitor", 0, rules, flag)

	if call.Err != nil {
		return nil, errors.Wrap(call.Err, "Failed to eavesdrop")
	}

	messages := make(chan *dbus.Message)

	conn.Eavesdrop(messages)

	return messages, nil
}
