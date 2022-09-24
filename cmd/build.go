package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/tatthien/giraffe/engine"
)

func Build() {
	start := time.Now()

	engine := engine.New()
	engine.ScanContent()
	engine.GenerateIndexPage()
	engine.GenerateTagIndexPage()
	engine.GenerateSingularPages()
	engine.GenerateTagPages()
	engine.GenerateRSS()
	engine.GeneratePostTypeArchive()
	engine.CopyStaticFiles()
	duration := time.Since(start)

	rows := map[string]int{
		"Pages":      len(engine.Posts),
		"Tags":       len(engine.Tags),
		"Post types": len(engine.PostTypes),
	}

	w := tabwriter.NewWriter(os.Stdout, 8, 8, 3, '\t', 0)
	fmt.Fprintln(w, "CONTENT\tTOTAL")
	for k, v := range rows {
		fmt.Fprintf(w, "%s\t%d\n", k, v)
	}
	fmt.Fprintf(w, "%s\t%s\n", "Build time", duration)
	w.Flush()
}
