package minesweeper

func isCellInside(r, c int, board *[]string) bool {
	return r >= 0 && r < len(*board) && c >= 0 && c < len((*board)[0])
}

// Annotate returns an annotated board
func Annotate(board []string) []string {

	for row := range board {
		cells := []rune(board[row])
		for col := range cells {
			if cells[col] == ' ' {
				bombs := 0
				neighbors := [][2]int{
					{row - 1, col - 1}, {row - 1, col}, {row - 1, col + 1},
					{row, col - 1}, {row, col + 1},
					{row + 1, col - 1}, {row + 1, col}, {row + 1, col + 1},
				}
				for _, neighbor := range neighbors {
					if isCellInside(neighbor[0], neighbor[1], &board) && board[neighbor[0]][neighbor[1]] == '*' {
						bombs++
					}
				}
				if bombs > 0 {
					cells[col] = rune(bombs + '0')
				}
			}
		}
		board[row] = string(cells)
	}

	return board
}
