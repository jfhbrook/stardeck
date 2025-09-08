package main

import (
	"flag"

	"github.com/jfhbrook/stardeck/service/lib"
)

func main() {
	flag.Parse()

	lib.Service()
}
