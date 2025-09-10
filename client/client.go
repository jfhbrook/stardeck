package client

import (
	"github.com/godbus/dbus/v5"
)

type StardeckClient struct {
	object dbus.BusObject
}

func NewStardeckClient(conn *dbus.Conn) *StardeckClient {
	obj := conn.Object("org.jfhbrook.stardeck", "/")
	client := StardeckClient {
		object: obj,
	}

	return &client
}

func (client *StardeckClient) SetWindow(name string) {
	client.object.Call("com.jfhbrook.stardeck.SetWindow", 0, name)
}
