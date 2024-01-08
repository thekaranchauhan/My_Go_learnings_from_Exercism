package perfect

import "errors"

var ErrOnlyPositive = errors.New("only positive numbers are allowed")

type Classification int

const (
	ClassificationInvalid Classification = iota
	ClassificationPerfect
	ClassificationAbundant
	ClassificationDeficient
)

// Classify returns the classification of the number.
func Classify(num int64) (Classification, error) {

	// Reject negative and zero numbers
	if num < 1 {
		return ClassificationInvalid, ErrOnlyPositive
	} else if num == 1 {
		// '1' is deficient
		return ClassificationDeficient, nil
	}

	// For each factor of num up to square root of num
	sum := int64(1)
	for divisor := int64(2); divisor <= num/divisor; divisor++ {
		if num%divisor == 0 {

			// Add the factor, and the complement factor if it's different
			sum += divisor
			if divisor != num/divisor {
				sum += num / divisor
			}

			// If sum exceeds num, return 'Abundant'
			if sum > num {
				return ClassificationAbundant, nil
			}
		}
	}

	// Return the apporpriate classification
	if sum == num {
		return ClassificationPerfect, nil
	} else {
		return ClassificationDeficient, nil
	}
}
