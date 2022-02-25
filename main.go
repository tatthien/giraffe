package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tatthien/giraffe/cmd"
)

const version = "v0.4.0"

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		cmd.Build()
		return
	}

	command := flag.Arg(0)
	switch command {
	case "serve":
		cmd.Serve()
	case "version":
		fmt.Println(version)
	default:
		log.Println("Unknown command:", command)
	}
}
