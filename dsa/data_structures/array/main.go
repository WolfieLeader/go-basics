package main

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/dsa/data_structures/array/dynamic"
	"github.com/WolfieLeader/go-basics/dsa/data_structures/array/matrix"
	"github.com/WolfieLeader/go-basics/dsa/data_structures/array/static"
)

func main() {
	fmt.Println("Static Array Example:")
	staticExample()
	fmt.Println()

	fmt.Println("Dynamic Array Example:")
	dynamicExample()
	fmt.Println()

	fmt.Println("Matrix Example:")
	matrixExample()
}

func staticExample() {
	s := static.New(1, 2, 3, 4, 5)
	fmt.Printf("- Find 2 with binary search (%v): %d\n", s, s.BinarySearch(2))
	s.Replace(2, 20, 5, 3)
	fmt.Printf("- Find 20 after replacing array (%v): %d\n", s, s.Search(20))
	fmt.Printf("- Reversed array (%v): %v\n", s, s.Reverse())
	s.Swap(4, 0)
	fmt.Printf("- Array after swapping 4 and 0: %v\n", s)
	s.Clear()
	c := s.Copy()
	c.Replace(5, 4, 3, 2, 1)
	fmt.Printf("- Are arrays equal? %v == %v: %t\n", s, c, s.Equal(c))
	s.Replace(5, 4, 3, 2, 1)
	fmt.Printf("- Are arrays equal? %v == %v: %t\n", s, c, s.Equal(c))
}

func dynamicExample() {
	d := dynamic.New(2, 4, 6, 8, 10, 12)
	fmt.Printf("- Find 4 with binary search (%v): %d\n", d, d.BinarySearch(4))
	d.Replace(10, 50, 20, 30, 40, 60)
	fmt.Printf("- Find 20 after replacing array (%v): %d\n", d, d.Search(20))
	fmt.Printf("- Reversed array (%v): %v\n", d, d.Reverse())
	d.Swap(1, 5)
	fmt.Printf("- Array after swapping 1 and 5: %v\n", d)
	d.Append(70, 80, 90)
	d.Prepend(0, -10, -20)
	b := d.Between(4, 7)
	fmt.Printf("- Array: %v, Between 4 and 7: %v\n", d, b)
	d.Clear()
	c := d.Copy()
	c.Replace(1, 2, 3, 4, 5)
	fmt.Printf("- Are arrays equal? %v == %v: %t\n", d, c, d.Equal(c))
	d.Replace(1, 2, 3, 4, 5)
	fmt.Printf("- Are arrays equal? %v == %v: %t\n", d, c, d.Equal(c))
}

func matrixExample() {
	m := matrix.New[int](3, 3)
	m.Replace([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	fmt.Printf("- Matrix after replacement:\n %v\n", m)
	e, _ := m.Get(1, 1)
	fmt.Printf("- Element at (1, 1): %d\n", e)
	fmt.Printf("- Row 0: %v\n", m.GetRow(0))
	fmt.Printf("- Column 1: %v\n", m.GetCol(1))
	m.Clear()
	c := m.Copy()
	c.Replace([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	fmt.Printf("- Are matrixes equal? %v == %v: %t\n", m, c, m.Equal(c))
	m.Replace([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	fmt.Printf("- Are matrixes equal? %v == %v: %t\n", m, c, m.Equal(c))
}
