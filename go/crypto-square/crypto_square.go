package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	pt = strings.ToLower(regexp.MustCompile(`\W`).ReplaceAllString(pt, ""))
	size := len(pt)

	c := int(math.Ceil(math.Sqrt(float64(size))))
	r := int(math.Ceil(float64(size) / float64(c)))

	pt += strings.Repeat(" ", c*r-size)
	result := ""

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			result += string(pt[i+j*c])
		}

		if i < c-1 {
			result += " "
		}
	}

	return result
}
