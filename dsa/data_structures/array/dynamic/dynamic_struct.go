package dynamic

import "cmp"

type DynamicArray[T cmp.Ordered] struct {
	data []T
}

func New[T cmp.Ordered](values ...T) *DynamicArray[T] {
	return &DynamicArray[T]{data: values}
}
