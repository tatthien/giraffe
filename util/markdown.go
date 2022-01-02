package util

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/spf13/viper"
	"github.com/tatthien/giraffe/model"
)

func IsMarkdownFile(path string) bool {
	re := regexp.MustCompile(`.md$`)
	return re.MatchString(path)
}

func GetFrontMatter(content string) (model.FrontMatter, string) {
	re := regexp.MustCompile(`(?m)^[\s\r\n]?---[\s\r\n]?$`)
	count := 1 // negative counter is equivalent to global case (replace all)

	content = re.ReplaceAllStringFunc(content, func(s string) string {
		if count == 0 {
			return s
		}

		count -= 1
		return re.ReplaceAllString(s, "")
	})

	// Splitting the front-matter and the body
	parts := strings.Split(content, "\n---\n")

	var fm model.FrontMatter
	var body string
	if len(parts) == 2 {
		body = parts[1]
	}

	yaml := []byte(parts[0])
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(yaml))

	fm.Title = viper.GetString("title")
	fm.Date = viper.GetTime("date")
	fm.Tags = viper.GetStringSlice("tags")

	return fm, body
}
