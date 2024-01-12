package pov

type Tree struct {
	value    string
	children []*Tree
}

func New(value string, children ...*Tree) *Tree {
	return &Tree{value, children}
}

func (t *Tree) Value() string {
	return t.value
}

func (t *Tree) Children() []*Tree {
	return t.children
}

func (t *Tree) tracePath(target string, trace []*Tree) []*Tree {
	trace = append(trace, t)

	if t.value == target {
		return trace
	}

	for _, kid := range t.children {
		if result := kid.tracePath(target, trace); result != nil {
			return result
		}
	}

	return nil
}

func (t *Tree) FromPov(targetValue string) *Tree {
	trace := t.tracePath(targetValue, make([]*Tree, 0, 4))

	if len(trace) == 0 {
		return nil
	}

	for i := len(trace) - 1; i > 0; i-- {
		curr, parent := trace[i], trace[i-1]

		for j, kid := range parent.children {
			if kid == curr {
				parent.children = append(parent.children[:j], parent.children[j+1:]...)
				break
			}
		}

		curr.children = append(curr.children, parent)
	}

	return trace[len(trace)-1]
}

func (t *Tree) PathTo(from, to string) []string {
	fromTree := t.FromPov(from)

	if fromTree == nil {
		return nil
	}

	trace := fromTree.tracePath(to, make([]*Tree, 0, 4))
	result := make([]string, 0, len(trace))

	for _, tree := range trace {
		result = append(result, tree.value)
	}

	return result
}
