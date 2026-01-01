package arrayqueue

import (
	"fmt"
	"strings"
)

func (q *ArrayQueue[T]) Enqueue(values ...T)  { q.data.Append(values...) }
func (q *ArrayQueue[T]) Dequeue() (T, bool)   { return q.data.Delete(0) }
func (q *ArrayQueue[T]) Peek() (T, bool)      { return q.data.Get(0) }
func (q *ArrayQueue[T]) Size() int            { return q.data.Length() }
func (q *ArrayQueue[T]) IsEmpty() bool        { return q.data.IsEmpty() }
func (q *ArrayQueue[T]) Clear()               { q.data.Clear() }
func (q *ArrayQueue[T]) ToSlice() []T         { return q.data.ToSlice() }
func (q *ArrayQueue[T]) Copy() *ArrayQueue[T] { return &ArrayQueue[T]{data: q.data.Copy()} }
func (q *ArrayQueue[T]) Equal(other *ArrayQueue[T]) bool {
	if q == other {
		return true
	}
	if q == nil || other == nil {
		return false
	}
	return q.data.Equal(other.data)
}
func (q *ArrayQueue[T]) String() string {
	if q.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < q.data.Length(); i++ {
		if i == 0 {
			sb.WriteString("(first) ")
		}
		v, _ := q.data.Get(i)
		fmt.Fprintf(&sb, "%v ", v)
		if i == q.data.Length()-1 {
			sb.WriteString("(last)")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
