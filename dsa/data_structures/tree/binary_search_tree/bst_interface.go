package bst

import "cmp"

type iterativeBinarySearchTree[T cmp.Ordered] interface {
	ContainsI(value T) bool
	InsertI(values ...T)
	DeleteI(values ...T) int

	TraverseInOrderI(fn func(value T))
	TraversePreOrderI(fn func(value T))
	TraversePostOrderI(fn func(value T))
	TraverseBreadthFirst(fn func(value T))

	MinI() (T, bool)
	MaxI() (T, bool)
	HeightI() int
	PredecessorI(value T) (T, bool)
	SuccessorI(value T) (T, bool)
	ToSliceI() []T
	EqualI(other *BinarySearchTree[T]) bool
	CopyI() *BinarySearchTree[T]
}

type recursiveBinarySearchTree[T cmp.Ordered] interface {
	ContainsR(value T) bool
	InsertR(values ...T)
	DeleteR(values ...T) int

	TraverseInOrderR(fn func(value T))
	TraversePreOrderR(fn func(value T))
	TraversePostOrderR(fn func(value T))

	MinR() (T, bool)
	MaxR() (T, bool)
	HeightR() int
	PredecessorR(value T) (T, bool)
	SuccessorR(value T) (T, bool)
	IsBalanced() bool
	ToSliceR() []T
	EqualR(other *BinarySearchTree[T]) bool
	CopyR() *BinarySearchTree[T]
	String() string
}

type binarySearchTree[T cmp.Ordered] interface {
	iterativeBinarySearchTree[T]
	recursiveBinarySearchTree[T]
	Size() int
	IsEmpty() bool
	Root() (T, bool)
	Clear()
}

var _ binarySearchTree[int] = (*BinarySearchTree[int])(nil)
