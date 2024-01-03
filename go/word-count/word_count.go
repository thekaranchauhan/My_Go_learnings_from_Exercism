package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`\w+('\w+)?`)

	count := make(Frequency)
	phrase = strings.ToLower(phrase)
	for _, v := range re.FindAllString(phrase, -1) {
		count[v]++
	}
	return count
}
