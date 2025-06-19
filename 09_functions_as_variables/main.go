package main

import "fmt"

func fibonacciWithLocalVar() func() int {
	// This is a closure that captures the local variables
	a, b := 0, 1

	// It returns a function that generates the next Fibonacci number
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	// This is a variable that holds an anonymous function
	someFunc := func(x int) int {
		return x * x
	}

	fmt.Println(someFunc(5)) // Output: 25
	fibFunc := fibonacciWithLocalVar()
	for range 10 {
		fmt.Println(fibFunc()) // Output: Fibonacci numbers
	}

	func() {
		fmt.Println("This is an anonymous function that runs immediately")
	}()
}
