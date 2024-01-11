package twobucket

import "errors"

type bucket struct {
	name, size, level int
}

var names = [...]string{"", "one", "two"}

func Solve(size1, size2, goal int, start string) (name string, moves, level int, err error) {
	// Error if bucket sizes are invalid or if no solution
	if size1 <= 0 || size2 <= 0 || goal <= 0 || (goal > size1 && goal > size2) ||
		(start != names[1] && start != names[2]) || goal%gcd(size1, size2) != 0 {
		return "", 0, 0, errors.New("invalid input or no solution")
	}

	// Create buckets
	this, other := &bucket{1, size1, 0}, &bucket{2, size2, 0}
	if start == names[2] {
		this, other = other, this
	}

	// Follow the algorithm until we reach the goal
	for moves = 0; this.level != goal && other.level != goal; moves++ {
		switch {
		case this.level == 0:
			this.level = this.size
		case other.size == goal:
			other.level = goal
		case other.level == other.size:
			other.level = 0
		default:
			pour := other.size - other.level
			if this.level < pour {
				pour = this.level
			}
			this.level, other.level = this.level-pour, other.level+pour
		}
	}

	// Return bucket name, num moves, and other level
	if other.level == goal {
		this, other = other, this
	}
	return names[this.name], moves, other.level, nil
}

func gcd(a, b int) int {
	for ; b != 0; a, b = b, a%b {
	}
	return a
}
