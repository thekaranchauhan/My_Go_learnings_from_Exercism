package matrix

import (
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix [][]int
type Pair [2]int

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	m := make(Matrix, 0, len(rows))
	for _, row := range rows {
		res := make([]int, 0, len(row)/2+1)
		for _, col := range strings.Split(row, " ") {
			val, _ := strconv.Atoi(col)
			res = append(res, val)
		}
		m = append(m, res)
	}
	return &m, nil
}

func (m *Matrix) Saddle() []Pair {
	result := []Pair{}
	if len(*m)*len((*m)[0]) == 1 {
		return result
	}
	for i := 0; i < len(*m); i++ {
		max := (*m)[i][0]
		for _, elem := range (*m)[i][1:] {
			if elem > max {
				max = elem
			}
		}
		pos := make([]int, 0, len((*m)[i]))
		for k, elem := range (*m)[i] {
			if elem == max {
				pos = append(pos, k)
			}
		}

		if len(pos) == 0 {
			continue
		}

		for _, l := range pos {
			min := 999999999999
			for j := 0; j < len(*m); j++ {
				if (*m)[j][l] < min {
					min = (*m)[j][l]
				}
			}

			if max == min {
				result = append(result, Pair{i + 1, l + 1})
			}
		}
	}
	return result
}
