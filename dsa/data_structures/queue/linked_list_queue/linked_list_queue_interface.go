package listqueue

import "cmp"

type linkedListQueue[T cmp.Ordered] interface {
	Enqueue(values ...T)
	Dequeue() (T, bool)
	Peek() (T, bool)

	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Equal(other *LinkedListQueue[T]) bool
	Copy() *LinkedListQueue[T]
	String() string
}

var _ linkedListQueue[int] = (*LinkedListQueue[int])(nil)
