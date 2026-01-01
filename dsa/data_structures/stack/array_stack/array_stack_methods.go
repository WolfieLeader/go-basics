package arraystack

import (
	"fmt"
	"strings"
)

func (s *ArrayStack[T]) Push(values ...T)     { s.data.Append(values...) }
func (s *ArrayStack[T]) Pop() (T, bool)       { return s.data.Delete(s.data.Length() - 1) }
func (s *ArrayStack[T]) Peek() (T, bool)      { return s.data.Get(s.data.Length() - 1) }
func (s *ArrayStack[T]) Size() int            { return s.data.Length() }
func (s *ArrayStack[T]) IsEmpty() bool        { return s.data.IsEmpty() }
func (s *ArrayStack[T]) Clear()               { s.data.Clear() }
func (s *ArrayStack[T]) ToSlice() []T         { return s.data.Reverse().ToSlice() }
func (s *ArrayStack[T]) Copy() *ArrayStack[T] { return &ArrayStack[T]{data: s.data.Copy()} }
func (s *ArrayStack[T]) Equal(other *ArrayStack[T]) bool {
	if s == other {
		return true
	}
	if s == nil || other == nil {
		return false
	}
	return s.data.Equal(other.data)
}
func (s *ArrayStack[T]) String() string {
	if s.IsEmpty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	for i := s.data.Length() - 1; i >= 0; i-- {
		if i == s.data.Length()-1 {
			sb.WriteString("(top) ")
		}
		v, _ := s.data.Get(i)
		fmt.Fprintf(&sb, "%v ", v)
		if i == 0 {
			sb.WriteString("(bottom)")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
