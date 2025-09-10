package main

import (
	"flag"

	"github.com/jfhbrook/stardeck/service"
)

func main() {
	flag.Parse()

	// TODO: Accept log level flag and config file location

	service.Service()
}
