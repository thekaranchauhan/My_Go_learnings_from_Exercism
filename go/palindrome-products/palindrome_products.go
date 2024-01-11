package palindrome

import (
	"errors"
	"math"
	"strconv"
)

// Define Product type here.
type Product struct {
	value          int
	Factorizations [][2]int
}

func isPalindrome(num int) bool {
	if num%10 == 0 {
		return false
	}
	if num < 10 {
		return true
	}
	if num < 100 {
		return num%11 == 0
	}

	rev := ""
	str := strconv.Itoa(num)

	for _, c := range str {
		rev = string(c) + rev
	}

	return rev == str
}

func Products(fmin, fmax int) (Product, Product, error) {
	palMin := Product{}
	palMax := Product{}

	if fmin > fmax {
		return palMin, palMax, errors.New("fmin > fmax")
	}

	var i, j int

	minFound := false
	maxFound := false

	for i = fmin; i <= fmax; i += 1 {
		if isPalindrome(i * i) {
			minFound = true
			palMin.value = i * i

			if fmin == 1 {
				palMin.Factorizations = append(palMin.Factorizations, [2]int{1, i})
			}

			palMin.Factorizations = append(palMin.Factorizations, [2]int{i, i})

			break
		}
	}

	if !minFound {
	minOut:
		for i = fmin; i <= fmax; i += 1 {
			for j = fmin; j <= fmax && i*j < palMin.value; j += 1 {
				prod := i * j

				if isPalindrome(prod) {
					minFound = true
					palMin.value = prod
					break minOut
				}
			}
		}

		if minFound {
			palMin.Factorizations = [][2]int{{int(math.Min(float64(i), float64(j))), int(math.Max(float64(i), float64(j)))}}
		}
	}

	for i = fmax; i >= fmin; i -= 1 {
		if isPalindrome(i * i) {
			maxFound = true
			palMax.value = i * i

			if fmin == 1 {
				palMax.Factorizations = append(palMax.Factorizations, [2]int{1, i * i})
			}

			palMax.Factorizations = append(palMax.Factorizations, [2]int{i, i})

			break
		}
	}

maxOut:
	for i = fmax; i >= fmin; i -= 1 {
		for j = fmax; (j >= fmin) && ((i * j) > palMax.value); j -= 1 {
			prod := i * j

			if isPalindrome(prod) {
				maxFound = true
				palMax.value = prod
				break maxOut
			}
		}
	}

	if i > 0 && j > 0 && maxFound {
		palMax.Factorizations = [][2]int{{int(math.Min(float64(i), float64(j))), int(math.Max(float64(i), float64(j)))}}
	}

	if !maxFound && !minFound {
		return palMin, palMax, errors.New("no palindromes")
	}

	return palMin, palMax, nil
}
