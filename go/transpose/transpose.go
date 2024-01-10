package transpose

import (
	"math"
	"strings"
)

func Transpose(input []string) []string {
	size := len(input)

	if size == 0 {
		return input
	}

	N := 0
	m := make([][]string, size)
	result := []string{}

	for i, row := range input {
		m[i] = strings.Split(row, "")
		N = int(math.Max(float64(N), float64(len(m[i]))))
	}

	for j := 0; j < N; j++ {
		var transpose strings.Builder

		for i := 0; i < size; i++ {
			if j < len(m[i]) {
				if i > transpose.Len() {
					transpose.WriteString(strings.Repeat(" ", i-transpose.Len()))
				}
				transpose.WriteByte(m[i][j][0])
			}
		}

		result = append(result, transpose.String())
	}

	result[len(result)-1] = strings.TrimRight(result[len(result)-1], " ")

	return result
}
