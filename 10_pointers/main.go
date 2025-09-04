package main

import (
	"errors"
	"fmt"
)

// This only modifies the local copy of x
func incorrectModify(x int) {
	x = 100
}

// This modifies the value at the address x points to
// The zero value of a pointer is nil, make sure to check for nil before doing anything with it!!!
func correctModify(x *int) error {
	if x == nil {
		return errors.New("nil pointer dereference")
	}

	*x = 100
	return nil
}

func main() {
	x := 42

	// `*T` is the pointer type for T, which means it holds the address of a variable of type T
	var p *int

	// `&` operator gets the address of a variable
	p = &x

	// `*` operator dereferences the pointer, getting the value at the address it points to
	fmt.Printf("Address of x(p): %p, Value at address of p: %d\n", p, *p)

	// You can also change the value of x through the pointer
	*p = 100
	fmt.Println("New value of x:", x)

	a := 23
	fmt.Println("Original value of a:", a)

	incorrectModify(a)
	fmt.Println("After incorrectModify, value of a:", a)

	correctModify(&a)
	fmt.Println("After correctModify, value of a:", a)

	// `new(T)` allocates memory for a variable of type T and returns a pointer to it
	pStr := new(string)

	*pStr = "Hello, World!"

	fmt.Println("Value at address pStr:", *pStr)
}
