package model

import (
	"fmt"
	"html/template"
	"time"
)

type Post struct {
	Title   string
	Date    time.Time
	Content template.HTML
	Slug    string
	Type    string
	Tags    []string
}

type ByDate []Post

type FrontMatter struct {
	Title string    `mapstructure:"title"`
	Date  time.Time `mapstructure:"date"`
	Tags  []string  `mapstructure:"tags"`
}

func (post *Post) Permarlink() string {
	return fmt.Sprintf("%s/%s.html", post.Type, post.Slug)
}

func (post *Post) FormattedDate(format string) string {
	if format == "" {
		format = "02-01-2006"
	}
	return post.Date.Format(format)
}

func (d ByDate) Len() int {
	return len(d)
}

func (d ByDate) Less(i, j int) bool {
	return d[i].Date.After(d[j].Date)
}

func (d ByDate) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
