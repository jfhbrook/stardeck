package service

import (
	"strconv"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/pkg/errors"
)

const (
	busName    = "org.jfhbrook.stardeck"
	objectPath = "/"
	ifaceName  = "org.jfhbrook.stardeck"
)

const intro = `
<node>
	<interface name="` + ifaceName + `">
		<method name="SetWindow">
			<arg direction="in" type="s"/>
		</method>
	</interface>` + introspect.IntrospectDataString + `
</node>`

type Iface struct {
	commands chan *command
}

func (i Iface) SetWindow(name string) *dbus.Error {
	i.commands <- makeSetWindowNameCommand(name)
	return nil
}

type DbusRequestNameError struct {
	Reply dbus.RequestNameReply
}

func (e DbusRequestNameError) Error() string {
	switch e.Reply {
	case dbus.RequestNameReplyPrimaryOwner:
		return "PrimaryOwner"
	case dbus.RequestNameReplyInQueue:
		return "InQueue"
	case dbus.RequestNameReplyExists:
		return "Exists"
	case dbus.RequestNameReplyAlreadyOwner:
		return "AlreadyOwner"
	}
	return strconv.FormatUint(uint64(e.Reply), 10)
}

func exportIface(conn *dbus.Conn) error {
	i := Iface{}

	conn.Export(i, objectPath, ifaceName)
	conn.Export(introspect.Introspectable(intro), objectPath, "org.freedesktop.DBus.Introspectable")

	reply, err := conn.RequestName(busName, dbus.NameFlagDoNotQueue)

	if err != nil {
		return errors.Wrap(err, "Failed to export DBus interface")
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		return errors.Wrap(DbusRequestNameError{Reply: reply}, "Failed to export DBus interface")
	}

	return nil
}
