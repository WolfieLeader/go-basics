package main

import (
	"fmt"
)

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
	panic("Panic!")
	fmt.Println("3 (won't be executed)")
}