package arraydeque

import "cmp"

type arrayDeque[T cmp.Ordered] interface {
	EnqueueFirst(values ...T)
	EnqueueLast(values ...T)
	DequeueFirst() (T, bool)
	DequeueLast() (T, bool)
	PeekFirst() (T, bool)
	PeekLast() (T, bool)

	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Equal(other *ArrayDeque[T]) bool
	Copy() *ArrayDeque[T]
	String() string
}

var _ arrayDeque[int] = (*ArrayDeque[int])(nil)
