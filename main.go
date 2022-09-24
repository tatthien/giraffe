package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tatthien/giraffe/cmd"
)

const version = "0.9.0"

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
		fmt.Printf("giraffe version: v%s\n", version)
	case "new":
		cmd.New(flag.Arg(1))
	case "help":
		cmd.Help()
	case "build":
		cmd.Build()
	default:
		log.Println("Unknown command:", command)
	}
}
