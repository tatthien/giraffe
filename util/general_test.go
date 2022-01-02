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
