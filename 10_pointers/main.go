package main

import (
	"errors"
	"fmt"
)

func pointExample() {
	fmt.Println("\nPointers Example:")

	var x int = 42 // Initialize an int variable

	// `*T` is the pointer type for T, which means it holds the address of a variable of type T
	var p *int // Declare a pointer to an int

	// `&` operator gets the address of a variable
	p = &x                          // Assign the address of x to p
	fmt.Println("Value of x:", x)   // Print the value of x
	fmt.Println("Address of x:", p) // Print the address of x

	// `*` operator dereferences the pointer, getting the value at the address it points to
	fmt.Println("Value at address p:", *p) // Print the value at the address p points to

	// You can also change the value of x through the pointer
	*p = 100                          // Change the value at the address p points to
	fmt.Println("New value of x:", x) // Print the new value of x
}

func incorrectModify(x int) {
	x = 100 // This only modifies the local copy of x
}

func correctModify(x *int) error {
	// The zero value of a pointer is nil, make sure to check for nil before doing anything with it!!!
	if x == nil {
		return errors.New("nil pointer dereference")
	}

	*x = 100 // This modifies the value at the address x points to
	return nil
}

func modifyExample() {
	fmt.Println("\nModify Example:")

	var a int = 42
	fmt.Println("Original value of a:", a)

	incorrectModify(a) // This will not change the value of a
	fmt.Println("After incorrectModify, value of a:", a)

	correctModify(&a) // Pass the address of a to modify it
	fmt.Println("After correctModify, value of a:", a)
}

func newExample() {
	fmt.Println("\nNew Example:")

	// `new(T)` allocates memory for a variable of type T and returns a pointer to it
	p := new(string)     // Allocate memory for a string and get a pointer to it
	*p = "Hello, World!" // Set the value at the address p points to

	fmt.Println("Value at address p:", *p) // Print the value at the address p points to
}

func main() {
	pointExample()
	modifyExample()
	newExample()
}
