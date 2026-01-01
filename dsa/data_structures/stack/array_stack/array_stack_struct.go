package arraystack

import (
	"cmp"

	"github.com/WolfieLeader/go-basics/dsa/data_structures/array/dynamic"
)

type ArrayStack[T cmp.Ordered] struct {
	data *dynamic.DynamicArray[T]
}

func New[T cmp.Ordered]() *ArrayStack[T] {
	return &ArrayStack[T]{data: dynamic.New[T]()}
}
