package bob

import (
	"strings"
)

// Hey returns Bob's response to a sentence.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	question := strings.HasSuffix(remark, "?")
	capitalized := strings.ToUpper(remark) != strings.
		ToLower(remark) && strings.ToUpper(remark) == remark

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case capitalized && question:
		return "Calm down, I know what I'm doing!"

	case capitalized:
		return "Whoa, chill out!"

	case question:
		return "Sure."

	default:
		return "Whatever."
	}
}
