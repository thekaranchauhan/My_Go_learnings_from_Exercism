package grains

import "errors"

// Square takes a Chess Board square number, and calculates the number of grains on it.
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("Invalid")
	}
	return 1 << uint64(square-1), nil
}

// Total calculaes the total number of grains on a chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
