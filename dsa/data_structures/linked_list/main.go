package main

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/dsa/data_structures/linked_list/doubly"
	"github.com/WolfieLeader/go-basics/dsa/data_structures/linked_list/singly"
)

func main() {
	fmt.Println("Singly Linked List Example:")
	singlyExample()
	fmt.Println()

	fmt.Println("Doubly Linked List Example:")
	doublyExample()
	fmt.Println()
}

func singlyExample() {
	s := singly.New[int]()
	fmt.Printf("- Empty: %v\n", s)
	s.AddFirst(3)
	fmt.Printf("- Element: %v\n", s)
	s.AddFirst(2, 1)
	s.AddLast(4, 5)
	fmt.Printf("- More: %v\n", s)
	s.DeleteAt(2)        // Delete 3
	n := s.GetNode(1)    // Get Node of 2
	s.DeleteAfterNode(n) // Delete 4
	fmt.Printf("- After Deletions: %v\n", s)
	fmt.Printf("- Reversed: %v\n", s.Reverse())
	s.Swap(2, 1)
	fmt.Printf("- Swap(2,1): %v\n", s)
	s.Clear()
	c := s.Copy()
	c.AddLast(1, 2, 3, 4)
	c.AddFirst(0)
	c = c.Reverse()
	fmt.Printf("- Are lists equal? %v == %v: %t\n", s, c, s.Equal(c))
	s.AddLast(4, 3, 2, 1, 0)
	fmt.Printf("- Are lists equal? %v == %v: %t\n", s, c, s.Equal(c))
}

func doublyExample() {
	d := doubly.New[int]()
	fmt.Printf("- Empty: %v\n", d)
	d.AddFirst(3)
	fmt.Printf("- Element: %v\n", d)
	d.AddFirst(2, 1)
	d.AddLast(4, 5)
	fmt.Printf("- More: %v\n", d)
	d.DeleteAt(2)        // Delete 3
	n := d.GetNode(1)    // Get Node of 2
	d.DeleteAfterNode(n) // Delete 4
	fmt.Printf("- After Deletions: %v\n", d)
	fmt.Printf("- Reversed: %v\n", d.Reverse())
	d.Swap(2, 1)
	fmt.Printf("- Swap(2,1): %v\n", d)
	d.Clear()
	c := d.Copy()
	c.AddLast(2, 4, 6, 8)
	c.AddFirst(0)
	c = c.Reverse()
	fmt.Printf("- Are lists equal? %v == %v: %t\n", d, c, d.Equal(c))
	d.AddLast(8, 6, 4, 2, 0)
	fmt.Printf("- Are lists equal? %v == %v: %t\n", d, c, d.Equal(c))
}
