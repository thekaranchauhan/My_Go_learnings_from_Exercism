package tree

import (
	"errors"
	"slices"
)

// Record is the input record format.
type Record struct{ ID, Parent int }

// Node is the output tree node format.
type Node struct {
	ID       int
	Children []*Node
}

// Build makes the tree
func Build(recs []Record) (*Node, error) {

	// Return nil reference if no input records
	if len(recs) == 0 {
		return nil, nil
	}

	// Sort the slice so children are processed after parents and in order
	// sort.Slice(recs, func(i, j int) bool { return recs[i].ID < recs[j].ID })
	slices.SortFunc(recs, func(a, b Record) int { return a.ID - b.ID })

	// Allocate all the space for the tree
	tree := make([]Node, len(recs))

	// For each input record
	for i, rec := range recs {

		// Validate record
		if rec.ID != i || (i != 0 && i <= rec.Parent) || (i == 0 && rec.Parent != 0) {
			return nil, errors.New("invalid input")
		}

		// Add node and parent's reference
		tree[i].ID = i
		if rec.ID > 0 {
			tree[rec.Parent].Children = append(tree[rec.Parent].Children, &tree[rec.ID])
		}
	}
	return &tree[0], nil
}
