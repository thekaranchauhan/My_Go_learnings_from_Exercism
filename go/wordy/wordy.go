package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	fullReg := regexp.MustCompile("What\\sis\\s-?\\d+(\\s(plus|minus|multiplied\\sby|divided\\sby)\\s-?\\d+)*\\?")
	if fullReg.MatchString(question) {
		numsReg := regexp.MustCompile("-?\\d+")
		operReg := regexp.MustCompile("(plus|minus|multiplied\\sby|divided\\sby)")

		operations := operReg.FindAllString(question, -1)
		numbers := numsReg.FindAllString(question, -1)

		result := strToInt(numbers[0])
		for i := 0; i < len(operations); i++ {
			switch operations[i] {
			case "plus":
				result += strToInt(numbers[i+1])
			case "minus":
				result -= strToInt(numbers[i+1])
			case "multiplied by":
				result *= strToInt(numbers[i+1])
			case "divided by":
				result /= strToInt(numbers[i+1])
			}
		}

		return result, true
	}
	return 0, false
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}
