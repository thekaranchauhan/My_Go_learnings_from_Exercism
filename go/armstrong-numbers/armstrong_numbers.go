package armstrong

import "math"

func IsNumber(n int) bool {
	digits := []int{}
	for i := n; i > 0; i /= 10 {
		digits = append(digits, i%10)
	}
	sum := 0
	for _, d := range digits {
		sum += int(math.Pow(float64(d), float64(len(digits))))
	}
	return sum == n
}
