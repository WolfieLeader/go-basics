package listdeque

import "cmp"

type linkedListDeque[T cmp.Ordered] interface {
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
	Equal(other *LinkedListDeque[T]) bool
	Copy() *LinkedListDeque[T]
	String() string
}

var _ linkedListDeque[int] = (*LinkedListDeque[int])(nil)
