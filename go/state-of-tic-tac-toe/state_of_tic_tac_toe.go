package stateoftictactoe

import (
	"errors"
	"strings"
)

type State string

const Win, Ongoing, Draw State = "win", "ongoing", "draw"

// StateOfTicTacToe returns the state of a game, given the board
func StateOfTicTacToe(board []string) (State, error) {
	b := board[0] + board[1] + board[2]
	turnsX, turnsO := strings.Count(b, "X"), strings.Count(b, "O")

	// Error if the number of turns is invalid or both players have a line of three
	if d := turnsX - turnsO; d < 0 || d > 1 || (isWinner(b, 'X') && isWinner(b, 'O')) {
		return "", errors.New("invalid state")
	}

	// Return win if a player has won
	if isWinner(b, 'X') || isWinner(b, 'O') {
		return Win, nil
	}

	// Return draw if X has had all 5 turns
	if turnsX+turnsO == 9 {
		return Draw, nil
	}

	// Else return ongoing
	return Ongoing, nil
}

// Check each line of three for a win
func isWinner(b string, p byte) bool {
	for _, l := range [...][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}} {
		if b[l[0]] == p && b[l[1]] == p && b[l[2]] == p {
			return true
		}
	}
	return false
}
