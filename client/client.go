package client

import (
	"github.com/godbus/dbus/v5"
)

type BusType int

const (
	Manage   bool = true
	NoManage      = false
)

type Client struct {
	object dbus.BusObject
}

func NewClient(conn *dbus.Conn) *Client {
	obj := conn.Object("org.jfhbrook.stardeck", "/")
	client := Client{
		object: obj,
	}

	return &client
}

func (client *Client) SetWindow(name string) error {
	return client.object.Call("org.jfhbrook.stardeck.SetWindow", 0, name).Store()
}

func (client *Client) SetLoopback(manage bool) error {
	return client.object.Call("org.jfhbrook.stardeck.SetLoopback", 0, manage).Store()
}
