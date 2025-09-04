package main

import (
	"fmt"
)

// Basic generic function to print `any`` type
func genericPrint[T any](value T) {
	fmt.Println(value)
}

// Generic function that can get any `comparable` type and check for equality
func isEqual[T comparable](a, b T) bool {
	return a == b
}

// Generic min function that has a generic union type constraint
// The `~` operator allows for type sets that include both the specified type and its underlying type
func min[T ~int | ~float64](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// Constraints interface for integer types
type Integers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Generic function to check if a number is even, constrained to integer types
func isEven[T Integers](n T) bool {
	return n%2 == 0
}

// Created a set type using generics
type Set[T comparable] map[T]struct{}

// Function to create a new Set with initial items
func newSet[T comparable](items ...T) Set[T] {
    set := make(Set[T])
    for _, item := range items {
        set[item] = struct{}{}
    }
    return set
}

// Implementing a Stringer interface for the Set type
func (s Set[T]) String() string {
    result := "{"
    for item := range s {
        result += fmt.Sprintf(" %v", item)
    }
    result += " }"
    return result
}

func main() {
    genericPrint("Hello, Generics!")
    genericPrint(42)

    fmt.Println("Is 5 equal to 3?", isEqual(5, 3))
    fmt.Println("Is Struct{A: 1} equal to Struct{A: 1}?", isEqual(struct{ A int }{A: 1}, struct{ A int }{A: 1}))

    fmt.Println("Min of 3 and 5:", min(3, 5))
    fmt.Println("Min of 3.5 and -2.1:", min(3.5, -2.1))

    fmt.Println("Is 4 even?", isEven(4))
    fmt.Println("Is 11 even?", isEven(11))

    mySet := newSet(1, 2, 3, 4, 5)
    fmt.Println("My Set:", mySet)
}
