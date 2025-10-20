package main

import "fmt"

func anonymousFunctionsExample() {
	// Immediately invoked function expression (IIFE)
	func() { fmt.Println("- This is an IIFE") }()

	squareFn := func(x int) int { return x * x }
	fmt.Println("Square of 5:", squareFn(5))
}

// Variadic functions allow you to pass a variable number of arguments
func variadicPrint(nums ...int) {
	if len(nums) == 0 {
		fmt.Println("- No numbers provided")
		return
	}

	fmt.Print("- Numbers:")
	for i, num := range nums {
		fmt.Printf(" [%d]: %d", i, num)
	}
	fmt.Println()
}

func variadicFunctionExample() {
	variadicPrint()        // No arguments
	variadicPrint(1, 2, 3) // Multiple arguments
}

// Returns a closure that generates Fibonacci numbers.
// It captures local variables a and b and updates them with each call.
func fibonacciClosure() func() int {
	a, b := 0, 1
	return func() int { a, b = b, a+b; return a }
}

func closureFunctionExample() {
	fibFn1 := fibonacciClosure()
	fibFn2 := fibonacciClosure()

	fmt.Println("- Fibonacci Sequence 1:")
	for range 5 {
		fmt.Printf("%d ", fibFn1())
	}
	fmt.Println()

	fmt.Println("- Fibonacci Sequence 2:")
	for range 10 {
		fmt.Printf("%d ", fibFn2())
	}
	fmt.Println()
}

// Recursion is a function calling itself with a smaller value until it reaches the base case.
// Base case: if n is 0 or 1, return 1.
// Note: For large n, this can lead to stack overflow.
func recursiveFactorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * recursiveFactorial(n-1)
}

func recursiveFactorialExample() {
	fmt.Println("- Factorial of 5:", recursiveFactorial(5))
	fmt.Println("- Factorial of 0:", recursiveFactorial(0))
}

// This is a higher-order function that takes another function as an argument (can be called a callback).
func doN(n int, fn func()) {
	for range n {
		fn()
	}
}

func higherOrderFunctionExample() {
	count := 0
	incrementFn := func() { count++ }
	doN(5, incrementFn)
	fmt.Println("- Count after incrementing 5 times:", count)
}
