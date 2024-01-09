package diamond

import (
	"bytes"
	"errors"
	"strings"
)

// Gen - returns a letter diamond
func Gen(input byte) (string, error) {

	// Validate input
	if input < 'A' || input > 'Z' {
		return "", errors.New("char out of range")
	}

	// Allocate string slice with all rows
	rowLen := 2*(input-'A') + 1
	rows := make([]string, rowLen)

	// For each char from 'A' to provided char
	for c := byte('A'); c <= input; c++ {

		//Build the row
		row := bytes.Repeat([]byte{' '}, int(rowLen))

		row[input-c], row[rowLen-1-input+c] = c, c
		//Add the row into the diamond, twice

		rows[c-'A'], rows[rowLen-1-c+'A'] = string(row), string(row)
	}

	// Return concatenated string
	return strings.Join(rows, "\n"), nil

}
