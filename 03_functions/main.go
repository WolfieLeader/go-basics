package main

import (
	"fmt"
	"math"
)

// In go function that starts with a capital letter is exported, meaning it can be accessed from other packages.
func printHelloWorld() {
	fmt.Println("Hello World!")
}

func add(a int, b int) int { return a + b }

func subtract(a, b int) (result int) {
	result = a - b
	return // Naked return, returns result
}

// Function can return multiple values
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 5, 2

	printHelloWorld()

	fmt.Println("- Addition:", add(a, b))
	fmt.Println("- Subtraction:", subtract(a, b))
	fmt.Println("- Multiplication:", multiply(a, b)) //Came from functions.go file

	x, y := swap(a, b)
	fmt.Printf("- Swapped from a=%d, b=%d to a=%d, b=%d\n", a, b, x, y)
	fmt.Println("- Power of 2^3:", math.Pow(2, 3)) // Using math package to calculate power
}
