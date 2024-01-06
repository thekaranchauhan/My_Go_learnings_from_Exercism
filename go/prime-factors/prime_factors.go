package prime

// Factors returns the prime factors of the given natural number.
func Factors(n int64) []int64 {
	r := make([]int64, 0)
	for i := int64(2); i <= n; i++ {
		for n%i == 0 {
			r = append(r, i)
			n /= i
		}
	}

	return r
}
