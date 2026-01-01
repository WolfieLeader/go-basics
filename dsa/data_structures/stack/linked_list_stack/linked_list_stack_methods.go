package liststack

import (
	"fmt"
	"strings"
)

func (s *LinkedListStack[T]) Push(values ...T) { s.data.AddFirst(values...) }
func (s *LinkedListStack[T]) Pop() (T, bool)   { return s.data.DeleteFirst() }
func (s *LinkedListStack[T]) Peek() (T, bool)  { return s.data.GetFirst() }
func (s *LinkedListStack[T]) Size() int        { return s.data.Size() }
func (s *LinkedListStack[T]) IsEmpty() bool    { return s.data.IsEmpty() }
func (s *LinkedListStack[T]) Clear()           { s.data.Clear() }
func (s *LinkedListStack[T]) ToSlice() []T     { return s.data.ToSlice() }
func (s *LinkedListStack[T]) Equal(other *LinkedListStack[T]) bool {
	if s == other {
		return true
	}
	if s == nil || other == nil {
		return false
	}
	return s.data.Equal(other.data)
}
func (s *LinkedListStack[T]) Copy() *LinkedListStack[T] {
	return &LinkedListStack[T]{data: s.data.Copy()}
}
func (s *LinkedListStack[T]) String() string {
	if s.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < s.data.Size(); i++ {
		if i == 0 {
			sb.WriteString("(top) ")
		}
		v, _ := s.data.Get(i)
		fmt.Fprintf(&sb, "%v ", v)
		if i == s.data.Size()-1 {
			sb.WriteString("(bottom)")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
