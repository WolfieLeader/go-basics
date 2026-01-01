package hashset

import (
	"fmt"
	"maps"
	"sort"
	"strings"
)

func (s *HashSet[T]) Size() int     { return len(s.data) }
func (s *HashSet[T]) IsEmpty() bool { return len(s.data) == 0 }
func (s *HashSet[T]) Clear()        { s.data = make(map[T]struct{}) }

func (s *HashSet[T]) Add(values ...T) {
	for _, v := range values {
		s.data[v] = struct{}{}
	}
}

func (s *HashSet[T]) Delete(values ...T) int {
	deletes := 0
	for _, v := range values {
		if _, ok := s.data[v]; ok {
			delete(s.data, v)
			deletes++
		}
	}
	return deletes
}

func (s *HashSet[T]) Contains(value T) bool {
	_, ok := s.data[value]
	return ok
}

func (s *HashSet[T]) ToSlice() []T {
	out := make([]T, 0, len(s.data))
	for k := range s.data {
		out = append(out, k)
	}
	return out
}

func (s *HashSet[T]) Equal(other *HashSet[T]) bool {
	if s == other {
		return true
	}
	if s == nil || other == nil {
		return false
	}
	return maps.Equal(s.data, other.data)
}

func (s *HashSet[T]) Traverse(fn func(value T) bool) {
	for k := range s.data {
		if !fn(k) {
			return
		}
	}
}

func (s *HashSet[T]) Copy() *HashSet[T] {
	out := New[T]()
	maps.Copy(out.data, s.data)
	return out
}

func (s *HashSet[T]) String() string {
	if s.Size() == 0 {
		return "[]"
	}

	values := make([]string, 0, len(s.data))
	for k := range s.data {
		values = append(values, fmt.Sprintf("%v", k))
	}

	sort.Strings(values)
	return "[" + strings.Join(values, " ") + "]"
}
