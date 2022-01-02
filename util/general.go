package util

import (
	"os"
	"regexp"
	"strings"
)

func Slugify(str string) string {
	regexps := map[string]string{
		"a": `a|á|à|ả|ã|ạ|ă|ắ|ằ|ẳ|ẵ|ặ|â|ấ|ầ|ẩ|ẫ|ậ`,
		"e": `e|é|è|ẽ|ẻ|ẹ|ê|ế|ề|ễ|ể|ệ`,
		"o": `o|ó|ò|õ|ỏ|ọ|ô|ố|ồ|ỗ|ổ|ộ|ơ|ớ|ờ|ỡ|ở|ợ`,
		"u": `u|ú|ù|ũ|ủ|ụ|ư|ứ|ừ|ữ|ử|ự`,
		"d": `đ`,
		"-": `\s`,
	}

	str = strings.ToLower(str)

	for repl, r := range regexps {
		re := regexp.MustCompile(r)
		str = re.ReplaceAllString(str, repl)
	}

	return str
}

func CreateDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
