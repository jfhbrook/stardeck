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

func Connect() (*Client, error) {
	conn, err := dbus.ConnectSessionBus()

	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to Stardeck service")
	}

	defer conn.Close()

	cl := NewClient(conn)

	return cl, nil
}

func (client *Client) SetWindow(name string) {
	client.object.Call("com.jfhbrook.stardeck.SetWindow", 0, name)
}

func (client *Client) SetLoopback(manage bool) {
	client.object.Call("com.jfhbrook.stardeck.SetLoopback", 0, manage)
}
