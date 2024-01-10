package linkedlist

import "errors"

type List struct {
	values []int
}

func New(values []int) *List {
	return &List{values}
}

func (l *List) Size() int {
	return len(l.values)
}

func (l *List) Push(element int) {
	l.values = append(l.values, element)
}

func (l *List) Pop() (int, error) {
	if len(l.values) == 0 {
		return 0, errors.New("can't pop an empty stack")
	}
	last := l.values[len(l.values)-1]
	l.values = l.values[:len(l.values)-1]
	return last, nil
}

func (l *List) Array() []int {
	return l.values
}

func (l *List) Reverse() *List {
	for i, j := 0, len(l.values)-1; i < j; i, j = i+1, j-1 {
		l.values[i], l.values[j] = l.values[j], l.values[i]
	}
	return l
}
