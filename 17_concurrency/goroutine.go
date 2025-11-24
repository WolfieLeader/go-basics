package main

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// They are used to perform concurrent (asynchronous) tasks in Go.
func goroutineExample() {
	// `go` keyword is used to start a new goroutine (created 2 goroutines here).
	go fmt.Println("- Foo")
	go func() { fmt.Println("- Bar") }()

	time.Sleep(100 * time.Millisecond) // Allow goroutines to finish before main exits.
	fmt.Println("Main goroutine finished.")
}

func goroutineWithoutWaitExample() {
	go fmt.Println("- Hello from goroutine")
	fmt.Println("Main goroutine finished (without waiting).")
}
