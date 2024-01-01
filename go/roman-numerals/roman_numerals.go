package romannumerals

import (
	"errors"
	"sort"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {

	if input < 1 || input > 3999 {
		return "", errors.New("Invalid input")
	}
	numbers := map[int]string{1: "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	output := ""

	keys := make([]int, 0)

	for k, _ := range numbers {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, i := range keys {
		if input/i > 0 {
			output += strings.Repeat(numbers[i], input/i)
			input -= (input / i) * i
		}

	}

	return output, nil
}
