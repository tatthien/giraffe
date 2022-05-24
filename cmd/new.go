package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/tatthien/giraffe/util"
)

func New(path string) {
	if path == "" {
		log.Fatalln("Missing path")
	}

	path = util.Slugify("content/" + path)
	util.CreateDir(filepath.Dir(path))

	markdownContent := fmt.Sprintf("---\ntitle: New Post\ndate: %s\ndraft: false\ntags:\n---\n", time.Now().Format("2006-01-02T15:04:05Z"))

	f, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	f.Write([]byte(markdownContent))

	fmt.Printf("New post created: %s\n", path)
}
