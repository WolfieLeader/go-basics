package singly

import "cmp"

type singlyLinkedList[T cmp.Ordered] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Copy() *SinglyLinkedList[T]

	AddFirst(values ...T)
	AddLast(values ...T)
	GetFirst() (T, bool)
	GetFirstNode() *Node[T]
	GetLast() (T, bool)
	GetLastNode() *Node[T]
	DeleteFirst() (T, bool)
	DeleteLast() (T, bool)

	SetAt(index int, value T) bool
	SetAtNode(node *Node[T], value T) bool
	InsertAt(index int, value T) bool
	InsertAtNode(node *Node[T], value T) bool
	InsertAfter(index int, value T) bool
	InsertAfterNode(node *Node[T], value T) bool

	DeleteAt(index int) (T, bool)
	DeleteAfter(index int) (T, bool)
	DeleteAtNode(node *Node[T]) (T, bool)
	DeleteAfterNode(node *Node[T]) (T, bool)
	DeleteValue(value T) bool

	Get(index int) (T, bool)
	GetNode(index int) *Node[T]
	ToSlice() []T
	Search(value T) int
	Contains(value T) bool

	Equal(other *SinglyLinkedList[T]) bool
	Traverse(fn func(index int, value T))
	Reverse() *SinglyLinkedList[T]
	IsSorted() bool
	Swap(index1, index2 int) bool
	String() string
}

var _ singlyLinkedList[int] = (*SinglyLinkedList[int])(nil)
