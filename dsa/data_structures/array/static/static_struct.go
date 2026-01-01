package static

import "cmp"

const SIZE = 5

type StaticArray[T cmp.Ordered] struct {
	data [SIZE]T
}

func New[T cmp.Ordered](values ...T) *StaticArray[T] {
	var data [SIZE]T

	if len(values) > SIZE {
		values = values[:SIZE] // Limit to SIZE elements
	}

	copy(data[:], values)
	return &StaticArray[T]{data: data}
}
