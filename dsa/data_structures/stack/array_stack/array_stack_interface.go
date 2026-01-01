package arraystack

import "cmp"

type arrayStack[T cmp.Ordered] interface {
	Push(values ...T)
	Pop() (T, bool)
	Peek() (T, bool)

	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Equal(other *ArrayStack[T]) bool
	Copy() *ArrayStack[T]
	String() string
}

var _ arrayStack[int] = (*ArrayStack[int])(nil)
