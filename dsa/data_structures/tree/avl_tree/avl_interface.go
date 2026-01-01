package avl

import "cmp"

type avlTree[T cmp.Ordered] interface {
	Contains(value T) bool
	Insert(values ...T) int
	Delete(values ...T) int
	Size() int
	IsEmpty() bool
	Clear()
	Root() (T, bool)
	Min() (T, bool)
	Max() (T, bool)
	Height() int
	BalanceFactor() int
	Successor(value T) (T, bool)
	Predecessor(value T) (T, bool)
	TraverseInOrder(fn func(value T))
	TraversePreOrder(fn func(value T))
	TraversePostOrder(fn func(value T))
	TraverseBreadthFirst(fn func(value T))
	IsBalanced() bool
	ToSlice() []T
	Equal(other *AVLTree[T]) bool
	Copy() *AVLTree[T]
	String() string
}

var _ avlTree[int] = (*AVLTree[int])(nil)
