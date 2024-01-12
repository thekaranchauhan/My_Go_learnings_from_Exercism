package ocr

import (
	"regexp"
	"strings"
)

var ref = map[string]string{
	"\n _ \n| |\n|_|\n   ": "0",
	"\n   \n  |\n  |\n   ": "1",
	"\n _ \n _|\n|_ \n   ": "2",
	"\n _ \n _|\n _|\n   ": "3",
	"\n   \n|_|\n  |\n   ": "4",
	"\n _ \n|_ \n _|\n   ": "5",
	"\n _ \n|_ \n|_|\n   ": "6",
	"\n _ \n  |\n  |\n   ": "7",
	"\n _ \n|_|\n|_|\n   ": "8",
	"\n _ \n|_|\n _|\n   ": "9",
}

func Recognize(input string) []string {
	digit, ok := ref[input]

	if ok {
		return []string{digit}
	}

	offset := 0
	grid := [][]string{{}}
	partPattern := regexp.MustCompile(`[\s_|]{3}`)

	for i, line := range strings.Split(input, "\n") {
		if i > 4 && i%4 == 1 {
			offset++
			grid = append(grid, []string{})
		}

		for index, part := range partPattern.FindAllStringSubmatch(line, -1) {
			if index >= len(grid[offset]) {
				grid[offset] = append(grid[offset], "")
			}
			grid[offset][index] += "\n" + strings.Join(part, "")
		}
	}

	result := make([]string, len(grid))

	for i, row := range grid {
		for _, raw := range row {
			val := "?"
			digit, ok := ref[raw]

			if ok {
				val = digit
			}

			result[i] += val
		}
	}

	return result
}

var recognizeDigit interface{}
