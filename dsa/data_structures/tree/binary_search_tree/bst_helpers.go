package bst

type stack[T any] struct{ data []T }
type queue[T any] struct{ data []T }

func newStack[T any](values ...T) *stack[T] { return &stack[T]{data: values} }
func newQueue[T any](values ...T) *queue[T] { return &queue[T]{data: values} }

func (s *stack[T]) Size() int     { return len(s.data) }
func (s *stack[T]) IsEmpty() bool { return len(s.data) == 0 }
func (q *queue[T]) Size() int     { return len(q.data) }
func (q *queue[T]) IsEmpty() bool { return len(q.data) == 0 }

func (s *stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *stack[T]) Pop() T {
	var zero T
	if len(s.data) == 0 {
		return zero
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *stack[T]) Peek() T {
	var zero T
	if len(s.data) == 0 {
		return zero
	}
	return s.data[len(s.data)-1]
}

func (q *queue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

func (q *queue[T]) Dequeue() T {
	var zero T
	if len(q.data) == 0 {
		return zero
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value
}

func (q *queue[T]) Peek() T {
	var zero T
	if len(q.data) == 0 {
		return zero
	}
	return q.data[0]
}
