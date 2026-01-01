package heap

import "cmp"

type Heap[T any] struct {
	data    []T
	aboveFn func(a, b T) bool
}

func New[T cmp.Ordered](isMinHeap bool) *Heap[T] {
	if isMinHeap {
		return &Heap[T]{aboveFn: func(a, b T) bool { return a < b }}
	}
	return &Heap[T]{aboveFn: func(a, b T) bool { return a > b }}
}

func NewByFunc[T any](aboveFn func(a, b T) bool) *Heap[T] {
	return &Heap[T]{aboveFn: aboveFn}
}
