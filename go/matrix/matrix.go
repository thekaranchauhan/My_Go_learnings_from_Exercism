package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

// New creates a new matrix with the given dimensions
func New(in string) (Matrix, error) {

	// Split the input into rows and allocate matrix space
	rows, err := strings.Split(in, "\n"), error(nil)
	m := make(Matrix, len(rows))

	// Parse, trim and validate each row
	for r, row := range rows {
		cells := strings.Fields(row)
		if r > 0 && len(cells) != len(m[0]) {
			return nil, errors.New("invalid column count")
		}

		// Allocate row space and parse each cell
		m[r] = make([]int, len(cells))
		for c, cell := range cells {
			if m[r][c], err = strconv.Atoi(cell); err != nil {
				return nil, err
			}
		}
	}
	return m, nil
}

// Rows returns the rows of the matrix
func (m *Matrix) Rows() [][]int {
	out := make([][]int, len(*m))
	for r, row := range *m {
		out[r] = append(make([]int, 0, len(*m)), row...)
	}
	return out
}

// Cols returns the columns of the matrix
func (m *Matrix) Cols() [][]int {
	out := make([][]int, len((*m)[0]))
	for r := range (*m)[0] {
		out[r] = make([]int, len(*m))
		for c := range *m {
			out[r][c] = (*m)[c][r]
		}
	}
	return out
}

// Set the value at the given row and column
func (m *Matrix) Set(r, c, v int) bool {
	if r < 0 || c < 0 || r >= len(*m) || c >= len((*m)[0]) {
		return false
	}
	(*m)[r][c] = v
	return true
}
