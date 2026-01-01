package main

import (
	"fmt"

	arraystack "github.com/WolfieLeader/go-basics/dsa/data_structures/stack/array_stack"
	liststack "github.com/WolfieLeader/go-basics/dsa/data_structures/stack/linked_list_stack"
)

func main() {
	fmt.Println("Array Stack Example:")
	arrayStackExample()
	fmt.Println()

	fmt.Println("Linked List Stack Example:")
	linkedListStackExample()
	fmt.Println()
}

func arrayStackExample() {
	s := arraystack.New[int]()
	s.Push(3, 2, 1)
	fmt.Printf("- Stack: %v\n", s)
	s.Pop()
	fmt.Printf("- Pop: %v\n", s)
	for i := range 10 {
		s.Push(i)
	}
	fmt.Printf("- 10 Pushes: %v\n", s)
	v, _ := s.Peek()
	fmt.Printf("- Top: %v\n", v)
	for range 5 {
		s.Pop()
	}
	fmt.Printf("- 5 Pops: %v\n", s)
	v, _ = s.Peek()
	fmt.Printf("- New Top: %v\n", v)
	s.Clear()
	c := s.Copy()
	c.Push(1, 2, 3, 4)
	fmt.Printf("- Are stacks equal? %v == %v: %t\n", s, c, s.Equal(c))
	s.Push(1, 2, 3, 4)
	fmt.Printf("- Are stacks equal? %v == %v: %t\n", s, c, s.Equal(c))
}

func linkedListStackExample() {
	s := liststack.New[int]()
	s.Push(3, 2, 1)
	fmt.Printf("- Stack: %v\n", s)
	s.Pop()
	fmt.Printf("- Pop: %v\n", s)
	for i := range 10 {
		s.Push(i)
	}
	fmt.Printf("- 10 Pushes: %v\n", s)
	v, _ := s.Peek()
	fmt.Printf("- Top: %v\n", v)
	for range 5 {
		s.Pop()
	}
	fmt.Printf("- 5 Pops: %v\n", s)
	v, _ = s.Peek()
	fmt.Printf("- New Top: %v\n", v)
	c := s.Copy()
	c.Push(1, 2, 3, 4)
	fmt.Printf("- Are stacks equal? %v == %v: %t\n", s, c, s.Equal(c))
	s.Push(1, 2, 3, 4)
	fmt.Printf("- Are stacks equal? %v == %v: %t\n", s, c, s.Equal(c))
}
