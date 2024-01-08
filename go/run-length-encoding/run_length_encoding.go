package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(in string) string {
	var out strings.Builder

	// For each char in string
	for i, count := 0, 1; i < len(in); i, count = i+1, count+1 {

		// If char is different to the next
		// output char count followed by char
		if i == len(in)-1 || in[i] != in[i+1] {
			if count > 1 {
				out.WriteString(strconv.Itoa(count))
			}
			out.WriteByte(in[i])
			count = 0
		}
	}
	return out.String()
}

func RunLengthDecode(input string) string {
	var out strings.Builder

	// While input left to process
	for len(input) > 0 {

		// Find the first non-digit
		charIndex, count := 0, 1
		for ; input[charIndex] >= '0' && input[charIndex] <= '9'; charIndex++ {
		}

		// Get char count
		if charIndex > 0 {
			count, _ = strconv.Atoi(string(input[:charIndex]))
		}

		// Add a string of count repeats of char
		// Then chop processed part of string for next iteration
		out.WriteString(strings.Repeat(string(input[charIndex]), count))
		input = input[charIndex+1:]
	}
	return out.String()
}
