package alphametics

import (
	"errors"
	"strings"
)

type char struct {
	val     int
	letter  byte
	leading bool
}

type alphametic struct {
	usedDigits [10]bool
	letters    [26]char
	index      []*char
	operands   []string
}

func Solve(puzzle string) (map[string]int, error) {
	a := alphametic{index: make([]*char, 0, 26)}
	a.operands = strings.FieldsFunc(puzzle, func(r rune) bool { return r == '+' || r == '=' || r == ' ' })

	for _, w := range a.operands {
		a.letters[w[0]-'A'].leading = true
		for _, c := range []byte(w) {
			a.letters[c-'A'].letter = c
		}
	}

	for i := range a.letters {
		if a.letters[i].letter > 0 {
			a.index = append(a.index, &a.letters[i])
		}
	}

	if !a.solveInner(0) {
		return nil, errors.New("no solution")
	}

	result := make(map[string]int, len(a.index))

	for _, v := range a.index {
		result[string(v.letter)] = v.val
	}

	return result, nil
}

func (a *alphametic) solveInner(depth int) bool {
	if depth == len(a.index) {
		tot := 0
		for _, o := range a.operands[:len(a.operands)-1] {
			tot += a.sumWord(o)
		}
		return tot == a.sumWord(a.operands[len(a.operands)-1])
	}

	char := a.index[depth]

	for digit := 0; digit < 10; digit++ {
		if (digit != 0 || !char.leading) && !a.usedDigits[digit] {
			a.usedDigits[digit], char.val = true, digit
			if a.solveInner(depth + 1) {
				return true
			}
			a.usedDigits[digit] = false
		}
	}

	return false
}

func (a *alphametic) sumWord(word string) (sum int) {
	for _, c := range []byte(word) {
		sum = 10*sum + a.letters[c-'A'].val
	}
	return sum
}
