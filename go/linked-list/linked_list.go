package linkedlist

import "fmt"

// ErrEmptyList is the error returned when the list is empty.
var ErrEmptyList = fmt.Errorf("empty list")

// Node is a double-linked list node. It holds a value, and references to next and previous nodes.
type Node struct {
	Value      any
	next, prev *Node
}

// Next returns the node's next node.
func (n *Node) Next() *Node { return n.next }

// Prev returns the node's previous node.
func (n *Node) Prev() *Node { return n.prev }

// List is a double linked list. It holds references to the first and last nodes.
type List struct{ first, last *Node }

// First returns the list's first node.
func (l *List) First() *Node { return l.first }

// Last returns the list's last node.
func (l *List) Last() *Node { return l.last }

// NewList creates and returns a new list from the supplied values
func NewList(values ...any) *List {
	l := &List{}
	for _, v := range values {
		l.Push(v)
	}
	return l
}

// Push pushes a value to the END of the list.
func (l *List) Push(v any) {
	newNode := Node{Value: v, prev: l.last}
	if l.first == nil {
		l.first = &newNode
	} else {
		l.last.next = &newNode
	}
	l.last = &newNode
}

// Unshift pushes a value to the START of the list.
func (l *List) Unshift(value any) {
	newNode := Node{Value: value, next: l.first}
	if l.first == nil {
		l.last = &newNode
	} else {
		l.first.prev = &newNode
	}
	l.first = &newNode
}

// Pop pops a value from the END of the list.
func (l *List) Pop() (any, error) {
	if l.first == nil {
		return nil, ErrEmptyList
	} else if l.last.prev != nil {
		// Clear node #2's 'next' link
		l.last.prev.next = nil
	} else {
		// Clear list if only one element
		l.first = nil
	}
	value := l.last.Value
	l.last = l.last.prev
	return value, nil
}

// Shift pops a value from the START of the list.
func (l *List) Shift() (any, error) {
	if l.first == nil {
		return nil, ErrEmptyList
	} else if l.first.next != nil {
		// clear node #2's 'prev' link
		l.first.next.prev = nil
	} else {
		// Clear list if only one element
		l.last = nil
	}
	value := l.first.Value
	l.first = l.first.next
	return value, nil
}

// Reverse reverses the order of all nodes in the list.
func (l *List) Reverse() {
	for n := l.first; n != nil; n = n.prev {
		n.prev, n.next = n.next, n.prev
	}
	l.first, l.last = l.last, l.first
}
