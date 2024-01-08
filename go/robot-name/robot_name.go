package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot represents a droid
type Robot struct{ name string }

// Create a shuffled pool of all possible names
var pool = func(size int) []int {
	p := make([]int, size)
	for i := 0; i < len(p); i++ {
		p[i] = i
	}
	rand.Shuffle(len(p), func(i, j int) { p[i], p[j] = p[j], p[i] })
	return p
}(26 * 26 * 10 * 10 * 10)

// Name returns a droid's name
func (r *Robot) Name() (string, error) {

	// Return name if already allocated
	if r.name != "" {
		return r.name, nil
	}

	// Return error if name pool exhausted
	if len(pool) == 0 {
		return "", errors.New("no more names available")
	}

	// Allocate next name from pool
	r.name = fmt.Sprintf("%c%c%03d", 'A'+pool[0]%26, 'A'+pool[0]/26%26, pool[0]/26/26)
	pool = pool[1:]
	return r.name, nil
}

// Reset resets a droid's name
func (r *Robot) Reset() { r.name = "" }
