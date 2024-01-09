package wordsearch

import "errors"

func Solve(words, puzzle []string) (map[string][2][2]int, error) {
	results := make(map[string][2][2]int, len(words))
	for _, word := range words {
		if !solveWord(word, puzzle, results) {
			return results, errors.New("word not found")
		}
	}
	return results, nil
}

func solveWord(word string, puzzle []string, results map[string][2][2]int) bool {
	for r, row := range puzzle {
		for c := range []byte(row) {
			if puzzle[r][c] == word[0] {
				for _, d := range [...]struct{ r, c int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, -1}, {1, 1}, {-1, -1}, {-1, 1}} {
					if er, ec, ok := solveWordDir(word[1:], puzzle, r+d.r, c+d.c, d.r, d.c); ok {
						results[word] = [2][2]int{{c, r}, {ec, er}}
						return true
					}
				}
			}
		}
	}
	return false
}

func solveWordDir(word string, puzzle []string, row, col, dr, dc int) (int, int, bool) {
	if row < 0 || col < 0 || row >= len(puzzle) || col >= len(puzzle[0]) || puzzle[row][col] != word[0] {
		return 0, 0, false
	} else if len(word) == 1 {
		return row, col, true
	}
	return solveWordDir(word[1:], puzzle, row+dr, col+dc, dr, dc)
}
