package arraydeque

import (
	"cmp"

	"github.com/WolfieLeader/go-basics/dsa/data_structures/array/dynamic"
)

type ArrayDeque[T cmp.Ordered] struct {
	data *dynamic.DynamicArray[T]
}

func New[T cmp.Ordered]() *ArrayDeque[T] {
	return &ArrayDeque[T]{data: dynamic.New[T]()}
}
