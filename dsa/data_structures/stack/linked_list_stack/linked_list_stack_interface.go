package liststack

import "cmp"

type linkedListStack[T cmp.Ordered] interface {
	Push(values ...T)
	Pop() (T, bool)
	Peek() (T, bool)

	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Equal(other *LinkedListStack[T]) bool
	Copy() *LinkedListStack[T]
	String() string
}

var _ linkedListStack[int] = (*LinkedListStack[int])(nil)
