package engine

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tatthien/giraffe/model"
	"github.com/tatthien/giraffe/util"
	"github.com/yuin/goldmark"
)

type AppEngine struct {
	SiteConfig  SiteConfig
	ContentDir  string
	DistDir     string
	TemplateDir string
	Posts       []model.Post
}

func New() *AppEngine {
	// Ensure dist directory is exists
	util.CreateDir("dist")

	// Load config
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	return &AppEngine{
		SiteConfig:  config,
		ContentDir:  "content",
		DistDir:     "dist",
		TemplateDir: "theme",
	}
}

func (engine *AppEngine) ScanContent() {
	var paths []string

	filepath.Walk(engine.ContentDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if util.IsMarkdownFile(path) {
			paths = append(paths, path)
		}

		return nil
	})

	for _, path := range paths {
		dir := filepath.Dir(path)
		folder := strings.Replace(dir, engine.ContentDir, "", 1)
		filename := filepath.Base(path)
		filename = strings.Replace(filename, ".md", "", 1)

		var post model.Post

		data, err := os.ReadFile(path)
		if err != nil {
			log.Println(err)
		}

		fm, body := util.GetFrontMatter(string(data))

		post.Type = strings.Replace(folder, "/", "", 1)
		post.Slug = util.Slugify(filename)
		post.Title = fm.Title
		post.Date = fm.Date
		post.Tags = fm.Tags

		var buf bytes.Buffer
		err = goldmark.Convert([]byte(body), &buf)
		if err != nil {
			log.Println(err)
		}
		post.Content = template.HTML(buf.String())

		engine.Posts = append(engine.Posts, post)
	}
}

func (engine *AppEngine) GenerateIndexPage() {
	sortedPosts := engine.Posts
	sort.Sort(model.ByDate(sortedPosts))
	data := struct {
		Site  SiteConfig
		Posts []model.Post
	}{engine.SiteConfig, sortedPosts}

	err := engine.SaveAsHTML("index.html", "index.html", data)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("generated: index.html")
	}
}

func (engine *AppEngine) GenerateSingluarPages() {
	for _, post := range engine.Posts {
		fileName := fmt.Sprintf("%s/%s.html", post.Type, post.Slug)
		err := engine.SaveAsHTML(fileName, "single.html", &post)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("generated: %s/%s.html\n", post.Type, post.Slug)
		}
	}
}

func (engine *AppEngine) SaveAsHTML(fileName, templateName string, data interface{}) error {
	tpl, err := template.ParseFiles(engine.TemplateDir + "/" + templateName)
	if err != nil {
		return err
	}

	fullPath := engine.DistDir + "/" + fileName

	err = util.CreateDir(filepath.Dir(fullPath))
	if err != nil {
		return err
	}

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tpl.Execute(f, data)
}

func (engine *AppEngine) CopyStaticFiles() {
	var paths []string

	filepath.Walk("theme/static", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	for _, path := range paths {
		dir := filepath.Dir(path)
		destDir := strings.Replace(dir, "theme/static", "dist", 1)
		filename := filepath.Base(path)

		// Create dir
		err := util.CreateDir(destDir)
		if err != nil {
			log.Println(err)
		}

		// Copy file
		srcFile, err := os.Open(path)
		if err != nil {
			log.Println(err)
		}
		defer srcFile.Close()

		destFile, err := os.Create(destDir + "/" + filename)
		if err != nil {
			log.Println(err)
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			log.Println(err)
		}
	}
}
