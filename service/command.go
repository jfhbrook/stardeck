package service

import (
	"github.com/rs/zerolog/log"
)

type commandType int

const (
	setWindowNameCommand commandType = 0
	setLoopbackCommand               = 1
)

type command struct {
	Type  commandType
	Value any
}

func makeSetWindowNameCommand(name string) *command {
	cmd := command{
		Type:  setWindowNameCommand,
		Value: name,
	}

	return &cmd
}

func makeSetLoopbackCommand(managed bool) *command {
	cmd := command{
		Type:  setLoopbackCommand,
		Value: managed,
	}

	return &cmd
}

func CommandRunner(commands chan *command) {
	for {
		log.Trace().Msg("Waiting for command")
		command := <-commands

		switch command.Type {
		case setWindowNameCommand:
			log.Debug().Any("name", command.Value).Msg("setWindowName")
		case setLoopbackCommand:
			log.Debug().Any("managed", command.Value).Msg("setLoopback")
		}
	}
}
