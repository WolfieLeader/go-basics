package main

import (
	"fmt"

	hashmap "github.com/WolfieLeader/go-basics/dsa/data_structures/hash_map/map"
	hashset "github.com/WolfieLeader/go-basics/dsa/data_structures/hash_map/set"
)

func main() {
	fmt.Println("HashMap Example")
	mapExample()
	fmt.Println()

	fmt.Println("HashSet Example")
	setExample()
	fmt.Println()
}

func mapExample() {
	m := hashmap.New[string, int]()
	m.Set("john", 25)
	m.Set("jane", 30)
	m.Set("doe", 22)
	m.Set("john", 28)
	fmt.Printf("- Map: %v\n", m)
	m.Delete("doe")
	fmt.Printf("- Map after deletion: %v\n", m)
	fmt.Printf("- Keys: %v\n", m.Keys())
	fmt.Printf("- Values: %v\n", m.Values())
	m.Clear()
	c := m.Copy()
	c.Set("alice", 35)
	c.Set("bob", 40)
	fmt.Printf("- Are maps equal? %v == %v: %t\n", m, c, m.Equal(c))
	m.Set("alice", 35)
	m.Set("bob", 40)
	fmt.Printf("- Are maps equal? %v == %v: %t\n", m, c, m.Equal(c))
}

func setExample() {
	s := hashset.New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(1)
	fmt.Printf("- Set: %v\n", s)
	s.Delete(2)
	fmt.Printf("- Set after removal: %v\n", s)
	fmt.Printf("- Contains 1: %v\n", s.Contains(1))
	fmt.Printf("- Contains 2: %v\n", s.Contains(2))
	s.Clear()
	c := s.Copy()
	c.Add(0, -1, -2, -3)
	fmt.Printf("- Are sets equal? %v == %v: %t\n", s, c, s.Equal(c))
	s.Add(0, -1, -2, -3)
	fmt.Printf("- Are sets equal? %v == %v: %t\n", s, c, s.Equal(c))
}
