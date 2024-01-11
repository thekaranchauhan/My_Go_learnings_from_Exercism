package binarysearchtree

type BinarySearchTree struct {
	data        int
	left, right *BinarySearchTree
}

func NewBst(i int) *BinarySearchTree { return &BinarySearchTree{data: i} }

func (b *BinarySearchTree) Insert(i int) *BinarySearchTree {
	if b == nil {
		return NewBst(i)
	}
	if i <= b.data {
		b.left = b.left.Insert(i)
	} else {
		b.right = b.right.Insert(i)
	}
	return b
}

func (b *BinarySearchTree) SortedData() []int {
	if b == nil {
		return nil
	}
	return append(append(b.left.SortedData(), b.data), b.right.SortedData()...)
}
