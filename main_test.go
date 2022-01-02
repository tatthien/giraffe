package main

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		actual := isMarkdownFile(tc[i].path)
		require.Equal(t, tc[i].expect, actual)
	}
}

func TestGenerateSlug(t *testing.T) {
	tc := []struct {
		str    string
		expect string
	}{
		{
			str:    "Xin chào",
			expect: "xin-chao",
		},
		{
			str:    "123-xin-chào",
			expect: "123-xin-chao",
		},
		{
			str:    "ạ-á-à-ã",
			expect: "a-a-a-a",
		},
	}

	for i := range tc {
		slug := generateSlug(tc[i].str)
		require.Equal(t, tc[i].expect, slug)
	}
}
