package service

import (
	"github.com/rs/zerolog/log"
)

type commandType int

const (
	setWindowNameCommand commandType = iota
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

func CommandRunner(commands chan *command) {
	for {
		command := <-commands

		switch command.Type {
		case setWindowNameCommand:
			log.Debug().Any("name", command.Value).Msg("SetWindowNameCommand")
		}
	}
}
