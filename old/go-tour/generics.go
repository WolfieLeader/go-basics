package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func TypeParamsExample() {
	sint := []int{0, 123, 456, 789}
	fmt.Println("Index of 456 in s:", Index(sint, 456))

	sstring := []string{"hello", "world", "foo", "bar"}
	fmt.Println("Index of 'foo' in s:", Index(sstring, "foo"))
}

func PrintAll[T any](s []T) {
	for i, v := range s {
		fmt.Printf("Index %d: %v\n", i, v)
	}
}

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

type List[T any] struct {
	val  T
	next *List[T]
}

func NewList[T any](val T) *List[T] {
	return &List[T]{val: val}
}

func (l *List[T]) Append(val T) *List[T] {
	newNode := &List[T]{val: val}
	if l == nil {
		return newNode
	}

	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return l
}

func (l *List[T]) Print() {
	current := l
	for current != nil {
		fmt.Printf("%v -> ", current.val)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *List[T]) Length() int {
	count := 0
	current := l
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func ListExample() {
	list := NewList(100) //? This is the head of the list
	list = list.Append(200)
	list = list.Append(300)
	list.Print()
	fmt.Printf("Length of the list: %d\n", list.Length())
}
