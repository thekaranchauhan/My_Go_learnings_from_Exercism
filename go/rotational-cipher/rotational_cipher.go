package rotationalcipher

import (
	"strings"
	"unicode"
)

var str = "abcdefghijklmnopqrstuvwxyz"

func RotationalCipher(s string, n int) string {
	var ss strings.Builder
	for _, v := range s {
		if !unicode.IsLetter(v) {
			ss.WriteRune(v)
			continue
		}

		if unicode.IsUpper(v) {
			str = strings.ToUpper(str)
		} else {
			str = strings.ToLower(str)
		}
		index := strings.Index(str, string(v))

		ss.WriteByte(str[(index+n)%26])
	}

	return ss.String()
}
