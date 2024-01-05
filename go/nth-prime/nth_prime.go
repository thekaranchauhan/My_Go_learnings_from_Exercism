package prime

import "errors"

func Nth(n int) (int, error) {
	c := 0
	if n <= 0 {
		return 0, errors.New("error")
	}
	for i := 2; ; i++ {
		if isPrime(i) {
			c++
			if c == n {
				return i, nil
			}
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
