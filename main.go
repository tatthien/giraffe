package main

import (
	"github.com/tatthien/giraffe/engine"
)

func main() {
	engine := engine.New()
	engine.ScanContent()

	engine.GenerateSingluarPages()
	engine.GenerateIndexPage()

	engine.CopyStaticFiles()
}
