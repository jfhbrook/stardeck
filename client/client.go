package client

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"
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
	call := client.object.Call("org.jfhbrook.stardeck.SetWindow", 0, name)

	if call.Err != nil {
		return errors.Wrap(call.Err, "Failed to set window title")
	}

	return nil
}

func (client *Client) SetLoopback(manage bool) error {
	call := client.object.Call("org.jfhbrook.stardeck.SetLoopback", 0, manage)

	if call.Err != nil {
		return errors.Wrap(call.Err, "Failed to set loopback")
	}

	return nil
}
