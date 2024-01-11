package dominoes

import "slices"

// Domino is a domino piece
type Domino [2]int

// MakeChain tries to create a domino chain from the provided bag of dominoes.
func MakeChain(bag []Domino) ([]Domino, bool) {
	return makeChain(slices.Clone(bag), make([]Domino, 0, len(bag)))
}

func makeChain(bag, chain []Domino) ([]Domino, bool) {
	// Return success if all dominoes used
	if len(chain) == len(bag) {
		return chain, true
	}

	// Try each domino in the bag
	for i, domino := range bag {
		// Rotate domino to match end of chain if required
		if len(chain) > 0 && chain[len(chain)-1][1] == domino[1] {
			domino[0], domino[1] = domino[1], domino[0]
		}

		// If domino not already used (zero marker) and matches target number
		if domino[0] != 0 && (len(chain) == 0 || chain[len(chain)-1][1] == domino[0]) {
			// Add domino to chain, zero mark used domino, try chain
			chain, bag[i][0] = append(chain, domino), 0
			if new, ok := makeChain(bag, chain); ok {
				return new, new[0][0] == new[len(new)-1][1]
			}
			chain, bag[i] = chain[:len(chain)-1], domino
		}
	}

	return chain, false
}
