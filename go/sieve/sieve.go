package sieve

func Sieve(limit int) []int {
	nums, results := make([]bool, limit+1), make([]int, 0, 256)
	for i := 2; i <= limit; i++ {
		if !nums[i] {
			results = append(results, i)
			for j := i * i; j <= limit; j += i {
				nums[j] = true
			}
		}
	}
	return results
}
