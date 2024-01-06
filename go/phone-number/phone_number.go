package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

func getDigits(s string) string {
	digits := []rune{}
	for _, char := range s {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		}
	}
	return string(digits)
}

func Number(phoneNumber string) (string, error) {
	numbers := getDigits(phoneNumber)
	if numbers[0] == '1' && len(numbers) == 11 {
		numbers = numbers[1:]
	}

	switch {
	case len(numbers) > 10:
		return "", errors.New("invalid number: contains more than ten digits")
	case len(numbers) < 10:
		return "", errors.New("invalid number: contains less than ten digits")
	case numbers[0] < '2':
		return "", errors.New("invalid area code")
	case numbers[3] < '2':
		return "", errors.New("invalid exchange code")
	default:
		return numbers, nil
	}
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
