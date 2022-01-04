package model

import "fmt"

type Tag struct {
	Name  string
	Slug  string
	Posts []Post
}

func (tag *Tag) Permalink() string {
	return fmt.Sprintf("/tags/%s.html", tag.Slug)
}

func (tag *Tag) AddPost(post Post) []Post {
	tag.Posts = append(tag.Posts, post)
	return tag.Posts
}

type Tags []Tag

func (tags Tags) Find(key string) Tag {
	for i := range tags {
		if tags[i].Name == key {
			return tags[i]
		}
	}

	return Tag{}
}
