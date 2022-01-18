package model

import (
	"fmt"
	"time"
)

const PostStatusDraft = "draft"
const PostStatusPublished = "published"

type Post struct {
	Title   string
	Date    time.Time
	Content string
	Slug    string
	Type    string
	Tags    []string
	Status  string
}

type Posts []Post

func (posts Posts) FindByTag(tag string) Posts {
	var foundPosts Posts
	for _, post := range posts {
		for _, t := range post.Tags {
			if t == tag {
				foundPosts = append(foundPosts, post)
			}
		}
	}

	return foundPosts
}

type ByDate []Post

type FrontMatter struct {
	Title string    `mapstructure:"title"`
	Date  time.Time `mapstructure:"date"`
	Tags  []string  `mapstructure:"tags"`
	Draft bool      `mapstructure:"draft"`
}

func (post *Post) Permarlink() string {
	return fmt.Sprintf("/%s/%s.html", post.Type, post.Slug)
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
