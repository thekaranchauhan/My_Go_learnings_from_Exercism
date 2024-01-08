package piglatin

import (
	"regexp"
	"strings"
)

func Sentence(sentence string) string {
	r := regexp.MustCompile(`^((?:[aeiou]|xr|yt).*|[^aeiou]*qu|[^aeiouy]+|[^aeiou]+)(.*)`)
	words := []string{}
	for _, w := range strings.Fields(strings.ToLower(sentence)) {
		words = append(words, r.ReplaceAllString(w, "$2$1")+"ay")
	}
	return strings.Join(words, " ")
}
