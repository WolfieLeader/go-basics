package main

import "fmt"

// Variadic functions allow you to pass a variable number of arguments
func printNumbers(nums ...int) {
	fmt.Println("Numbers:", nums)
}

// fibonacciGenerator returns a closure that generates Fibonacci numbers.
// It captures local variables a and b and updates them with each call.
func fibonacciGenerator() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// recursiveFactorial calculates the factorial of n recursively.
// Recursion is a function calling itself with a smaller value until it reaches the base case.
// Base case: if n is 0 or 1, return 1.
func recursiveFactorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}

func main() {
	// Anonymous function assigned to a variable
	square := func(x int) int {
		return x * x
	}
	fmt.Println("Square of 5:", square(5)) // Output: 25

	// Variadic function example
	printNumbers(3, 6, 9, 12, 15)

	// Closure-based Fibonacci generator
	fib := fibonacciGenerator()
	fmt.Println("First 10 Fibonacci numbers:")
	for range 10 {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()

	// Immediately invoked anonymous function
	func() {
		fmt.Println("This is an anonymous function that runs immediately.")
	}()

	// Factorial using recursion
	fmt.Printf("Factorial of %d: %d\n", 5, recursiveFactorial(5))
}
