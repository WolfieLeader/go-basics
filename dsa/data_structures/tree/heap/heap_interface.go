package heap

type heap[T any] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Peek() (T, bool)
	Pop() (T, bool)
	Push(values ...T)
	Contains(value T) bool
	Heapify(values ...T)
	ToSlice() []T
	Copy() *Heap[T]
	TraverseBreadthFirst(fn func(value T))
	Equal(other *Heap[T]) bool
	String() string
}

var _ heap[int] = (*Heap[int])(nil)
