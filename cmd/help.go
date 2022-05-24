package cmd

import "fmt"

func Help() {
	// Print help
	fmt.Println("Giraffe: An opinionated static site generator")
	fmt.Println("")
	fmt.Println("Usage: giraffe [command] [arguments]")
	fmt.Println("")
	fmt.Println("Available commands:")
	fmt.Println("  serve        Serve the site")
	fmt.Println("  new [path]   Create new content for your site")
	fmt.Println("  version      Print the version number of Giraffe")
	fmt.Println("")
}
