package cmd

import "fmt"

func Help() {
	fmt.Println(`
Giraffe: An opinionated static site generator

Usage: giraffe [command] [arguments]

Available commands:
	serve		start a server, watch files changed and rebuild
	build		generate static files
	new [path]	create a new markdown file
	version		print the cli version
	`)
}
