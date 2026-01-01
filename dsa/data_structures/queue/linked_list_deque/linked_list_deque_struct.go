package listdeque

import (
	"cmp"

	"github.com/WolfieLeader/go-basics/dsa/data_structures/linked_list/doubly"
)

type LinkedListDeque[T cmp.Ordered] struct {
	data *doubly.DoublyLinkedList[T]
}

func New[T cmp.Ordered]() *LinkedListDeque[T] {
	return &LinkedListDeque[T]{data: doubly.New[T]()}
}
