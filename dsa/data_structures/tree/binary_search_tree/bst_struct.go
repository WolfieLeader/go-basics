package bst

import "cmp"

type Node[T cmp.Ordered] struct {
	Value T
	left  *Node[T]
	right *Node[T]
}

type BinarySearchTree[T cmp.Ordered] struct {
	root *Node[T]
	size int
}

func New[T cmp.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}
