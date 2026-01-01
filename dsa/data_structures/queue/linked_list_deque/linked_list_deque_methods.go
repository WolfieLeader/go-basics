package listdeque

import (
	"fmt"
	"strings"
)

func (d *LinkedListDeque[T]) EnqueueFirst(values ...T) { d.data.AddFirst(values...) }
func (d *LinkedListDeque[T]) EnqueueLast(values ...T)  { d.data.AddLast(values...) }
func (d *LinkedListDeque[T]) DequeueFirst() (T, bool)  { return d.data.DeleteFirst() }
func (d *LinkedListDeque[T]) DequeueLast() (T, bool)   { return d.data.DeleteLast() }
func (d *LinkedListDeque[T]) PeekFirst() (T, bool)     { return d.data.GetFirst() }
func (d *LinkedListDeque[T]) PeekLast() (T, bool)      { return d.data.GetLast() }
func (d *LinkedListDeque[T]) Size() int                { return d.data.Size() }
func (d *LinkedListDeque[T]) IsEmpty() bool            { return d.data.IsEmpty() }
func (d *LinkedListDeque[T]) Clear()                   { d.data.Clear() }
func (d *LinkedListDeque[T]) ToSlice() []T             { return d.data.ToSlice() }
func (d *LinkedListDeque[T]) Equal(other *LinkedListDeque[T]) bool {
	if d == other {
		return true
	}
	if d == nil || other == nil {
		return false
	}
	return d.data.Equal(other.data)
}
func (d *LinkedListDeque[T]) Copy() *LinkedListDeque[T] {
	return &LinkedListDeque[T]{data: d.data.Copy()}
}
func (d *LinkedListDeque[T]) String() string {
	if d.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < d.data.Size(); i++ {
		if i == 0 {
			sb.WriteString("(first) ")
		}
		v, _ := d.data.Get(i)
		fmt.Fprintf(&sb, "%v ", v)
		if i == d.data.Size()-1 {
			sb.WriteString("(last)")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
