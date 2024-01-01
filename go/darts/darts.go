package darts

import (
	"math"
)

func Score(x_position, y_position float64) int {
	toss_coordinates := math.Sqrt(x_position*x_position + y_position*y_position)
	switch {
	case toss_coordinates <= 1.0:
		return 10
	case toss_coordinates <= 5.0:
		return 5
	case toss_coordinates <= 10.0:
		return 1
	default:
		return 0
	}
}
