package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(s string) bool {
	s = strings.NewReplacer("-", "").Replace(s)
	if len(s) != 10 {
		return false
	}
	res := 0

	for i, v := range s {
		if i < 9 && unicode.IsLetter(v) {
			return false
		}

		dig := int(v - '0')
		if v == 'X' {
			dig = 10
		}
		res += dig * (10 - i)
	}
	return res%11 == 0
}
