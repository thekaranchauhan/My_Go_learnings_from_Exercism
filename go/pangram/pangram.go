package pangram

import "unicode"

func IsPangram(input string) bool {
	check := [26]bool{}
	for _, v := range input {
		if unicode.IsLetter(v) {
			index := unicode.ToLower(v) - 'a'
			check[index] = true
		}
	}
	for _, v := range check {
		if !v {
			return false
		}
	}
	return true
}
