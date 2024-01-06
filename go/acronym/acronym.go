package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	result := ""
	list := regexp.MustCompile(`[A-Za-z']+`).FindAllString(s, -1)
	for _, word := range list {
		result += strings.ToUpper(string(word[0]))
	}
	return result
}
