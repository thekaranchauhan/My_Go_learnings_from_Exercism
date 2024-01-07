package bottlesong

import "strings"

var cnt = []string{"No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
var s = map[bool]string{true: "s"}

func Recite(start, verses int) (out []string) {
	for i := start; i > start-verses; i-- {
		out = append(out, "",
			cnt[i]+" green bottle"+s[i != 1]+" hanging on the wall,",
			cnt[i]+" green bottle"+s[i != 1]+" hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be "+strings.ToLower(cnt[i-1])+" green bottle"+s[i-1 != 1]+" hanging on the wall.")
	}
	return out[1:]
}
