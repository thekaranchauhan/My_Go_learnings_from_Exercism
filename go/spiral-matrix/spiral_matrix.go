package spiralmatrix

// SpiralMatrix returns a square matrix of numbers in spiral order.
func SpiralMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	x, y, dir := 0, 0, 0
	for count := 1; count <= n*n; count++ {
		res[x][y] = count
		u := x + dx[dir]
		v := y + dy[dir]
		if u < 0 || u >= n || v < 0 || v >= n || res[u][v] != 0 {
			dir = (dir + 1) % 4
		}
		x, y = x+dx[dir], y+dy[dir]
	}
	return res
}
