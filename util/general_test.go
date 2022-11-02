package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
		slug := Slugify(tc[i].str)
		require.Equal(t, tc[i].expect, slug)
	}
}

func TestSliceContains(t *testing.T) {
	tc := []struct {
		slice  []string
		str    string
		expect bool
	}{
		{
			[]string{"a", "b", "c"},
			"a",
			true,
		},
		{
			[]string{"a", "b", "c"},
			"d",
			false,
		},
		{
			[]string{},
			"",
			false,
		},
	}

	for i := range tc {
		ok := SliceContains(tc[i].slice, tc[i].str)
		require.Equal(t, tc[i].expect, ok)
	}
}
