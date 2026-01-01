package arraydeque

import (
	"fmt"
	"strings"
)

func (d *ArrayDeque[T]) EnqueueFirst(values ...T) { d.data.Prepend(values...) }
func (d *ArrayDeque[T]) EnqueueLast(values ...T)  { d.data.Append(values...) }
func (d *ArrayDeque[T]) DequeueFirst() (T, bool)  { return d.data.Delete(0) }
func (d *ArrayDeque[T]) DequeueLast() (T, bool)   { return d.data.Delete(d.data.Length() - 1) }
func (d *ArrayDeque[T]) PeekFirst() (T, bool)     { return d.data.Get(0) }
func (d *ArrayDeque[T]) PeekLast() (T, bool)      { return d.data.Get(d.data.Length() - 1) }
func (d *ArrayDeque[T]) Size() int                { return d.data.Length() }
func (d *ArrayDeque[T]) IsEmpty() bool            { return d.data.IsEmpty() }
func (d *ArrayDeque[T]) Clear()                   { d.data.Clear() }
func (d *ArrayDeque[T]) ToSlice() []T             { return d.data.ToSlice() }
func (d *ArrayDeque[T]) Equal(other *ArrayDeque[T]) bool {
	if d == other {
		return true
	}
	if d == nil || other == nil {
		return false
	}
	return d.data.Equal(other.data)
}
func (d *ArrayDeque[T]) Copy() *ArrayDeque[T] { return &ArrayDeque[T]{data: d.data.Copy()} }
func (d *ArrayDeque[T]) String() string {
	if d.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < d.data.Length(); i++ {
		if i == 0 {
			sb.WriteString("(first) ")
		}
		v, _ := d.data.Get(i)
		fmt.Fprintf(&sb, "%v ", v)
		if i == d.data.Length()-1 {
			sb.WriteString("(last)")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
