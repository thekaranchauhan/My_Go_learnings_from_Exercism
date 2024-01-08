package allyourbase

import "errors"

func ConvertToBase(baseIn int, digits []int, baseOut int) (output []int, err error) {

	// Bases must be 2 or higher
	if baseIn < 2 {
		return nil, errors.New("input base must be >= 2")
	} else if baseOut < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	// Convert input to int
	var sum int
	for _, d := range digits {
		if sum = sum*baseIn + d; d < 0 || d >= baseIn {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	// Convert int to output array
	for ; sum > 0 || len(output) == 0; sum /= baseOut {
		output = append([]int{(sum % baseOut)}, output...)
	}
	return output, nil
}
