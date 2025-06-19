package main

import "fmt"

func main() {
	// This is a variable that holds an anonymous function
	someFunc := func(x int) int {
		return x * x
	}

	fmt.Println(someFunc(5)) // Output: 25
}
