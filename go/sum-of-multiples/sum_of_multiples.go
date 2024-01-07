package summultiples

func SumMultiples(limit int, divisors ...int) int {
	multiples := make([]bool, limit)
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := divisor; i < limit; i += divisor {
			multiples[i] = true
		}
	}
	var sum int
	for i, b := range multiples {
		if b {
			sum += i
		}
	}
	return sum
}
