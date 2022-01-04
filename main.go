package main

import "github.com/tatthien/giraffe/engine"

func main() {
	engine := engine.New()

	engine.ScanContent()

	engine.GenerateIndexPage()
	engine.GenerateTagIndexPage()
	engine.GenerateSingluarPages()
	engine.GenerateTagPages()

	engine.CopyStaticFiles()
}
