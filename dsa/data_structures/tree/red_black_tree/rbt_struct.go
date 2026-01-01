package redblacktree

import "cmp"

type Node[T cmp.Ordered] struct {
	Value  T
	left   *Node[T]
	right  *Node[T]
	parent *Node[T]
	isRed  bool
}

type RedBlackTree[T cmp.Ordered] struct {
	root *Node[T]
	size int
}

func New[T cmp.Ordered]() *RedBlackTree[T] {
	return &RedBlackTree[T]{}
}
