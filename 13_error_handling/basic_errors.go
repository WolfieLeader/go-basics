package main

import (
	"errors"
	"fmt"
)

// Common pattern for error handling in Go
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// When a function starts with "must", it indicates that it will panic if an error occurs
func mustDivide(a, b int) int {
	result, err := divide(a, b)
	if err != nil {
		panic(err) // This will stop the program and print the error
	}
	return result
}

func deferCount(){
	// The defer keyword means that the function will be executed after the surrounding function returns
	defer func() { fmt.Println("5 (defer)") }()

	// Defer function that recovers from panic, it MUST be defined before the panic occurs
	// Deferred function calls are executed in LIFO (Last In, First Out) order
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		fmt.Println("4 (defer)")
	}()
	fmt.Println("1")
	fmt.Println("2")
	// This is a panic, it will cause the program to stop execution
	panic("random panic message")
	fmt.Println("3 (won't be executed)")
}

func basicErrorsExample() {
	fmt.Println("\nBasic Errors Example:")
	
	deferCount()

	result, err := divide(10, 0)
	// This is how you handle errors in Go
	// If err is not nil, it means an error occurred
	if err != nil {
		fmt.Println("Tried to divide by zero")
		// Usually you would return, but here we want to continue
	} else {
		fmt.Printf("10 / 0 = %d\n", result)
	}

	result, err = divide(9, 3)
	if err != nil {
		fmt.Println("Tried to divide by zero")
	} else {
		fmt.Printf("9 / 3 = %d\n", result)
	}

	
	result = mustDivide(8, 2)
	fmt.Printf("8 / 2 = %d\n", result)
	
	defer func() {
		if r:=recover(); r!=nil {
			fmt.Println("Recovered from panic in main:", r)
		}
		}()
		result = mustDivide(8, 0)
		fmt.Printf("8 / 0 = %d\n", result) // This will not be executed due to panic
}