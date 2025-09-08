package lib

import (
	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/rs/zerolog/log"
)

const (
	busName    = "org.jfhbrook.stardeck"
	objectPath = "/"
	ifaceName  = "org.jfhbrook.stardeck"
)

const intro = `
<node>
	<interface name="` + ifaceName + `">
		<method name="setWindow">
			<arg direction="in" type="s"/>
		</method>
	</interface>` + introspect.IntrospectDataString + `
</node>`

type Iface struct{}

func (i Iface) setWindow(name string) *dbus.Error {
	log.Debug().Any("name", name).Msg("Set window name")
	return nil
}

type DbusRequestNameError struct {
	Reply dbus.RequestNameReply
}

func (e DbusRequestNameError) Error() string {
	return string(e.Reply)
}

func exportIface(conn *dbus.Conn) error {
	i := Iface{}

	conn.Export(i, objectPath, ifaceName)
	conn.Export(introspect.Introspectable(intro), objectPath, "org.freedesktop.DBus.Introspectable")

	reply, err := conn.RequestName(busName, dbus.NameFlagDoNotQueue)

	if err != nil {
		return err
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		return DbusRequestNameError{Reply: reply}
	}

	return nil
}
