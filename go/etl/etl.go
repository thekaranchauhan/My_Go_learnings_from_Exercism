package etl

import (
	"strings"
)

type legacy map[int][]string

func Transform(given legacy) (returned map[string]int) {
	returned = make(map[string]int)

	for key, values := range given {
		for _, val := range values {
			returned[strings.ToLower(val)] = key
		}
	}
	return
}
