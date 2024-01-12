package say

import (
	"strconv"
	"strings"
)

var (
	ref = map[int64]string{
		0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
		6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
		11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
		15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen",
		19: "nineteen", 20: "twenty", 30: "thirty", 40: "forty",
		50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety",
	}
	power = map[int64]string{
		3: "hundred", 4: "thousand", 7: "million", 10: "billion",
	}
	powers = []int{3, 4, 7, 10}
)

func Say(n int64) (string, bool) {
	if n < 0 || n >= 1e12 {
		return "", false
	}
	if _, ok := ref[n]; ok {
		return ref[n], true
	}

	result := []string{}
	source := strings.Split(strconv.Itoa(int(n)), "")

	for i := 0; i < len(source); i++ {
		if len(source)-i < 3 {
			trail := source[len(source)-2:]
			key, _ := strconv.Atoi(strings.Join(trail, ""))

			if _, ok := ref[int64(key)]; ok {
				if key != 0 {
					result = append(result, ref[int64(key)])
				}
			} else {
				num1, _ := strconv.Atoi(trail[0])
				num2, _ := strconv.Atoi(trail[1])

				result = append(result, ref[int64(num1*10)]+"-"+ref[int64(num2)])
			}

			break
		}

		j := len(powers) - 1

		for ; (j >= 0) && (powers[j] > len(source)-i); j-- {
		}

		gap := len(source) - i - powers[j] + 1
		sub, _ := strconv.Atoi(strings.Join(source[i:i+gap], ""))

		if sub != 0 {
			res, _ := Say(int64(sub))
			result = append(result, res, power[int64(powers[j])])
		}

		i += gap - 1
	}

	return strings.Join(result, " "), true
}
