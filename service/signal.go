package service

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/crystalfontz"
)

func signalHandler(client *crystalfontz.Client) {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals

	err := client.Splash()

	if err != nil {
		log.Warn().Msg(err.Error())
	}

	err = client.StoreBootState(crystalfontz.NilFloat, crystalfontz.NilInt)

	if err != nil {
		log.Warn().Msg(err.Error())
	}
}
