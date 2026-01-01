package static

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/dsa/searching_algorithms/searching"
)

func (a *StaticArray[T]) Length() int { return len(a.data) }
func (a *StaticArray[T]) Clear()      { *a = StaticArray[T]{} }

func (a *StaticArray[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(a.data) {
		var zero T
		return zero, false
	}

	return a.data[index], true
}

func (a *StaticArray[T]) Set(index int, value T) bool {
	if index < 0 || index >= len(a.data) {
		return false
	}

	a.data[index] = value
	return true
}

func (a *StaticArray[T]) Replace(values ...T) {
	var zero T
	a.Fill(zero)

	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}
	copy(a.data[:], values)
}

func (a *StaticArray[T]) Fill(value T) {
	for i := range a.data {
		a.data[i] = value
	}
}

func (a *StaticArray[T]) IsSorted() bool {
	for i := 1; i < len(a.data); i++ {
		if a.data[i-1] > a.data[i] {
			return false
		}
	}
	return true
}

func (a *StaticArray[T]) Search(value T) int {
	return searching.LinearSearch(a.data[:], value)
}

func (a *StaticArray[T]) BinarySearch(value T) int {
	if !a.IsSorted() {
		return -1 // Binary search requires sorted array
	}
	return searching.BinarySearch(a.data[:], value)
}

func (a *StaticArray[T]) Contains(value T) bool {
	return a.Search(value) != -1
}

func (a *StaticArray[T]) Traverse(fn func(index int, value T)) {
	for index, value := range a.data {
		fn(index, value)
	}
}

func (a *StaticArray[T]) Swap(index1, index2 int) bool {
	length := len(a.data)
	if index1 < 0 || index1 >= length || index2 < 0 || index2 >= length {
		return false
	}

	if index1 == index2 {
		return true
	}

	a.data[index1], a.data[index2] = a.data[index2], a.data[index1]
	return true
}

func (a *StaticArray[T]) ToSlice() []T {
	out := make([]T, len(a.data))
	copy(out, a.data[:])
	return out
}

func (a *StaticArray[T]) Equal(other *StaticArray[T]) bool {
	if a == other {
		return true
	}
	if a == nil || other == nil {
		return false
	}
	return a.data == other.data
}

func (a *StaticArray[T]) Copy() *StaticArray[T] {
	var out StaticArray[T]
	copy(out.data[:], a.data[:])
	return &out
}

func (a *StaticArray[T]) Reverse() *StaticArray[T] {
	out := a.data
	left, right := 0, len(out)-1
	for left < right {
		out[left], out[right] = out[right], out[left]
		left++
		right--
	}
	return &StaticArray[T]{data: out}
}

func (a *StaticArray[T]) String() string {
	return fmt.Sprintf("%v", a.data)
}
