package main

import (
	"github.com/rs/zerolog/log"
)

type CommandType int

const (
	SetWindowNameCommand CommandType = iota
)

type Command struct {
	Type  CommandType
	Value any
}

func CommandRunner(commands chan *Command) {
	for {
		command := <-commands

		switch command.Type {
		case SetWindowNameCommand:
			log.Debug().Any("name", command.Value).Msg("SetWindowNameCommand")
		}
	}
}
