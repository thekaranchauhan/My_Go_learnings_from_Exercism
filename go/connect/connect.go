package connect

const WHITE, BLACK = 'O', 'X'

func ResultOf(board []string) (string, error) {
	for c := range board[0] {
		if follow(board, c, 0, WHITE) {
			return string(WHITE), nil
		}
	}

	for r := range board {
		if follow(board, 0, r, BLACK) {
			return string(BLACK), nil
		}
	}

	return "", nil
}

func follow(board []string, c, r int, colour byte) bool {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[r]) || board[r][c] != colour {
		return false
	}

	if (colour == WHITE && r == len(board)-1) || (colour == BLACK && c == len(board[0])-1) {
		return true
	}

	row := []byte(board[r])
	row[c] = '*'
	board[r] = string(row)

	directions := [][2]int{{0, -1}, {1, -1}, {1, 0}, {0, 1}, {-1, 1}, {-1, 0}}

	for _, d := range directions {
		if follow(board, c+d[0], r+d[1], colour) {
			return true
		}
	}

	return false
}
