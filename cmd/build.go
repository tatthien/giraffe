package cmd

import (
	"fmt"
	"time"

	"github.com/tatthien/giraffe/engine"
)

func Build() {
	start := time.Now()
	fmt.Println("Start building site...")
	fmt.Println("")

	engine := engine.New()
	engine.ScanContent()
	engine.GenerateIndexPage()
	engine.GenerateTagIndexPage()
	engine.GenerateSingluarPages()
	engine.GenerateTagPages()
	engine.GenerateRSS()
	engine.GeneratePostTypeArchive()
	engine.CopyStaticFiles()

	fmt.Println("  Content        | Total")
	fmt.Println("-----------------+--------")
	fmt.Printf("  Pages          | %d\n", len(engine.Posts))
	fmt.Printf("  Tags           | %d\n", len(engine.Tags))
	fmt.Printf("  Post types     | %d\n", len(engine.PostTypes))

	duration := time.Since(start)
	fmt.Println("")
	fmt.Printf("Build time %s\n", duration)
}
