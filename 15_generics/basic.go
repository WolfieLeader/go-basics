package main

import "fmt"

// Basic generic function to print `any`` type
func genericPrint[T any](value T) {
	fmt.Println(value)
}

// Generic function that can get any `comparable` type and check for equality
func isEqual[T comparable](a, b T) bool {
	return a == b
}

