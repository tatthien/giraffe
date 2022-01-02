package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tatthien/giraffe/model"
)

func TestIsMarkdownFile(t *testing.T) {
	tc := []struct {
		expect bool
		path   string
	}{
		{
			expect: true,
			path:   "foo.md",
		},
		{
			expect: true,
			path:   "123.md",
		},
		{
			expect: true,
			path:   "foo-123.md",
		},
		{
			expect: false,
			path:   "foo",
		},
		{
			expect: false,
			path:   "foo.txt",
		},
	}

	for i := range tc {
		actual := IsMarkdownFile(tc[i].path)
		require.Equal(t, tc[i].expect, actual)
	}
}

func TestGetFrontMatter(t *testing.T) {
	tc := []struct {
		expect  model.FrontMatter
		content string
	}{
		{
			expect: model.FrontMatter{
				Title: "Hello",
				Date:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				Tags:  []string{"foo", "bar"},
			},
			content: `
---
title: Hello
tags: foo,bar
date: 2022-01-01T00:00:00.00Z
---

Hello world!
			`,
		},
	}

	for i := range tc {
		fmt.Println(tc[i].content)
		actual, _ := GetFrontMatter(tc[i].content)
		require.Equal(t, tc[i].expect.Title, actual.Title)
		require.Equal(t, tc[i].expect.Tags, actual.Tags)
		require.WithinDuration(t, tc[i].expect.Date, actual.Date, time.Second)
	}
}
