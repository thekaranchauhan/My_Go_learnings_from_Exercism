package anagram

import (
	"strings"
)

func Check(subject string, candidate string) bool {
	subject = strings.ToUpper(subject)
	candidate = strings.ToUpper(candidate)
	if subject == candidate {
		return false
	}
	candidateMap := make(map[rune]int)
	for _, v := range candidate {
		candidateMap[v]++
	}
	for _, v := range subject {
		_, ok := candidateMap[v]
		if !ok {
			return false
		}
		candidateMap[v]--
		if candidateMap[v] == 0 {
			delete(candidateMap, v)
		}
	}
	return len(candidateMap) == 0
}
func Detect(subject string, candidates []string) []string {
	var result []string
	for _, v := range candidates {
		if Check(subject, v) {
			result = append(result, v)
		}
	}
	return result
}
