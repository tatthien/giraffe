package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func Serve() {
	// Build content
	Build()

	// Start the server
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/", http.FileServer(http.Dir("dist")))

		server := http.Server{
			Addr:    ":3333",
			Handler: mux,
		}

		fmt.Println("Serving on http://localhost:3333")
		log.Fatal(server.ListenAndServe())
	}()

	// Watch files changes
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(".", watchDir); err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("modified file:", event.Name)
				Build()
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	<-done
}

func watchDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		if strings.HasPrefix(path, "theme") || strings.HasPrefix(path, "content") {
			return watcher.Add(path)
		}
	}

	return nil
}
