package listqueue

import (
	"fmt"
	"strings"
)

func (q *LinkedListQueue[T]) Enqueue(values ...T) { q.data.AddLast(values...) }
func (q *LinkedListQueue[T]) Dequeue() (T, bool)  { return q.data.DeleteFirst() }
func (q *LinkedListQueue[T]) Peek() (T, bool)     { return q.data.GetFirst() }
func (q *LinkedListQueue[T]) Size() int           { return q.data.Size() }
func (q *LinkedListQueue[T]) IsEmpty() bool       { return q.data.IsEmpty() }
func (q *LinkedListQueue[T]) Clear()              { q.data.Clear() }
func (q *LinkedListQueue[T]) ToSlice() []T        { return q.data.ToSlice() }
func (q *LinkedListQueue[T]) Equal(other *LinkedListQueue[T]) bool {
	if q == other {
		return true
	}
	if q == nil || other == nil {
		return false
	}
	return q.data.Equal(other.data)
}
func (q *LinkedListQueue[T]) Copy() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{data: q.data.Copy()}
}
func (q *LinkedListQueue[T]) String() string {
	if q.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < q.data.Size(); i++ {
		if i == 0 {
			sb.WriteString("(first) ")
		}
		v, _ := q.data.Get(i)
		fmt.Fprintf(&sb, "%v ", v)
		if i == q.data.Size()-1 {
			sb.WriteString("(last)")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
