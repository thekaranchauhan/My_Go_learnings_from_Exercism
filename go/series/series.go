package series

func All(n int, s string) []string {
	ans := []string{}
	for i := 0; i <= len(s)-n; i++ {
		ans = append(ans, s[i:i+n])
	}
	return ans
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
