package arrayqueue

import "cmp"

type arrayQueue[T cmp.Ordered] interface {
	Enqueue(values ...T)
	Dequeue() (T, bool)
	Peek() (T, bool)

	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Equal(other *ArrayQueue[T]) bool
	Copy() *ArrayQueue[T]
	String() string
}

var _ arrayQueue[int] = (*ArrayQueue[int])(nil)
