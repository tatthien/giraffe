package engine

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tatthien/giraffe/markdown"
	"github.com/tatthien/giraffe/model"
	"github.com/tatthien/giraffe/util"
)

type AppEngine struct {
	SiteConfig  SiteConfig
	ContentDir  string
	OutputDir   string
	TemplateDir string
	Posts       model.Posts
	Tags        model.Tags
	PostTypes   []string
}

func New() *AppEngine {

	// Load config
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	// Ensure dist directory is exists
	util.CreateDir(config.OutputDir)

	return &AppEngine{
		SiteConfig:  config,
		ContentDir:  config.ContentDir,
		OutputDir:   config.OutputDir,
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

		if fm.Draft {
			post.Status = model.PostStatusDraft
		} else {
			post.Status = model.PostStatusPublished
		}

		// Collect all post types
		if !util.SliceContains(engine.PostTypes, post.Type) {
			engine.PostTypes = append(engine.PostTypes, post.Type)
		}

		post.Content, err = markdown.Render(body)
		if err != nil {
			log.Println(err)
		}

		if post.Status == model.PostStatusPublished {
			engine.Posts = append(engine.Posts, post)
		}
	}

	// Retrieve tags
	for _, post := range engine.Posts {
		for _, t := range post.Tags {
			tag := engine.Tags.Find(t)

			if tag.Name == "" {
				tag.Name = t
				tag.Slug = util.Slugify(t)

				engine.Tags = append(engine.Tags, tag)
			}
		}
	}
}

func (engine *AppEngine) GenerateIndexPage() {
	sortedPosts := engine.Posts
	sort.Sort(model.ByDate(sortedPosts))

	data := map[string]interface{}{
		"Posts":  sortedPosts,
		"IsHome": true,
	}

	err := engine.SaveAsHTML("index.html", "index.html", data)
	if err != nil {
		log.Println(err)
	}
}

func (engine *AppEngine) GenerateSingularPages() {
	for _, post := range engine.Posts {
		fileName := fmt.Sprintf("%s/%s.html", post.Type, post.Slug)
		data := map[string]interface{}{
			"Post":       post,
			"IsSingular": true,
		}
		err := engine.SaveAsHTML(fileName, "single.html", data)
		if err != nil {
			log.Println(err)
		}
	}
}

func (engine *AppEngine) GenerateTagIndexPage() {
	sortedTags := engine.Tags
	sort.Sort(model.TagsByName(sortedTags))

	// Generate tags.html
	err := engine.SaveAsHTML("tags/index.html", "tags.html", map[string]interface{}{
		"Tags": sortedTags,
	})
	if err != nil {
		log.Println(err)
	}
}

func (engine *AppEngine) GenerateTagPages() {
	for _, tag := range engine.Tags {
		tag.Posts = engine.Posts.FindByTag(tag.Name)
		fileName := fmt.Sprintf("tags/%s.html", tag.Slug)
		data := map[string]interface{}{
			"Tag":       tag,
			"IsArchive": true,
		}
		err := engine.SaveAsHTML(fileName, "tag.html", data)
		if err != nil {
			log.Println(err)
		}
	}
}

func (engine *AppEngine) GenerateRSS() {
	sortedPosts := engine.Posts
	sort.Sort(model.ByDate(sortedPosts))
	data := map[string]interface{}{
		"Posts": sortedPosts,
		"Site":  engine.SiteConfig,
	}

	rssTemplate := `
	{{ $baseURL := .Site.BaseURL }}
	<rss version="2.0">
	<channel>
		<title>{{ .Site.Title }}</title>
		<link>{{ $baseURL }}</link>
		<description>{{ .Site.Description }}</description>
		{{ range .Posts }}
		{{ if eq .Type "posts"}}
		<item>
			<title>{{ .Title }}</title>
			<link>{{ $baseURL }}{{ .Permarlink }}</link>
			<description></description>
			<pubDate>{{ .Date.Format "2006-01-02T15:04:05" }}</pubDate>
		</item>
		{{ end }}
		{{ end }}
	</channel>
	</rss>
	`

	t, err := template.New("").Parse(rssTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(engine.OutputDir + "/rss.xml")
	if err != nil {
		log.Fatalln(err)
	}

	t.Execute(f, data)
}

func (engine *AppEngine) GeneratePostTypeArchive() {
	for _, postType := range engine.PostTypes {
		var posts model.Posts
		for _, post := range engine.Posts {
			if post.Type == postType {
				posts = append(posts, post)
			}
		}

		data := map[string]interface{}{
			"Posts": posts,
		}

		templateName := postType + ".html"
		if _, err := os.Stat("theme/" + templateName); os.IsNotExist(err) {
			templateName = "post-type.html"
		}

		err := engine.SaveAsHTML(postType+"/index.html", templateName, data)
		if err != nil {
			log.Println(err)
		}
	}
}

func (engine *AppEngine) SaveAsHTML(fileName, templateName string, data map[string]interface{}) error {
	tpl := compileTemplate(templateName)

	fullPath := engine.OutputDir + "/" + fileName

	err := util.CreateDir(filepath.Dir(fullPath))
	if err != nil {
		return err
	}

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Add site config as global variable
	data["Site"] = map[string]interface{}{
		"BaseURL":     engine.SiteConfig.BaseURL,
		"Title":       engine.SiteConfig.Title,
		"Description": engine.SiteConfig.Description,
		"Posts":       engine.Posts,
		"Tags":        engine.Tags,
	}

	return tpl.ExecuteTemplate(f, templateName, data)
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

func compileTemplate(templateName string) *template.Template {
	t := template.New("")

	funcMap := template.FuncMap{
		"safe_html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"join": func(a []string, sep string) string {
			return strings.Join(a, sep)
		},
		"slugify": util.Slugify,
	}

	t = template.Must(t.Funcs(funcMap).ParseGlob("theme/common/*.html"))

	return template.Must(t.ParseFiles("theme/" + templateName))
}
