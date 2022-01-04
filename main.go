package main

import (
	"fmt"

	"github.com/tatthien/giraffe/engine"
)

func main() {
	engine := engine.New()

	engine.ScanContent()

	engine.GenerateIndexPage()
	engine.GenerateTagIndexPage()
	engine.GenerateSingluarPages()
	engine.GenerateTagPages()

	engine.CopyStaticFiles()

	fmt.Printf("Generated: %d posts\n", len(engine.Posts))
	fmt.Printf("Generated: %d tags\n", len(engine.Tags))
}
