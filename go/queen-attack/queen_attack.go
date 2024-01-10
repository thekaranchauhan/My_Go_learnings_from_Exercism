package queenattack

import (
	"fmt"
	"regexp"
)

func CanQueenAttack(whiteQ, blackQ string) (bool, error) {
	re := regexp.MustCompile("([a-h])([1-8])")
	if whiteQ == blackQ || !re.MatchString(whiteQ) || !re.MatchString(blackQ) {
		return false, fmt.Errorf("invalid position")
	}

	xDiff, yDiff := int(whiteQ[0])-int(blackQ[0]), int(whiteQ[1])-int(blackQ[1])

	return xDiff == 0 || yDiff == 0 || xDiff == yDiff || xDiff == -yDiff, nil
}
