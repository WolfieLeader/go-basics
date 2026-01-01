package avl

type queue[T any] struct{ data []T }

func newQueue[T any](values ...T) *queue[T] { return &queue[T]{data: values} }

func (q *queue[T]) Size() int     { return len(q.data) }
func (q *queue[T]) IsEmpty() bool { return len(q.data) == 0 }

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
