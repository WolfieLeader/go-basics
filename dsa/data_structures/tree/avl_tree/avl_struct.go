package avl

import "cmp"

type Node[T cmp.Ordered] struct {
	Value  T
	left   *Node[T]
	right  *Node[T]
	height int
}

type AVLTree[T cmp.Ordered] struct {
	root *Node[T]
	size int
}

func New[T cmp.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{}
}
