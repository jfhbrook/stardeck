package main

import (
	"flag"

	"github.com/jfhbrook/stardeck/service"
)

func main() {
	flag.Parse()

	service.Service()
}
