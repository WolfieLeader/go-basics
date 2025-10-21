package basic

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// They are used to perform concurrent (asynchronous) tasks in Go.
func GoroutineExample() {
	fmt.Println("\nGoroutine Example:")

	// `go` keyword is used to start a new goroutine.
	go fmt.Println("Foo")
	go func() { fmt.Println("Bar") }()
	go sayBaz()

	// Sleeping to allow goroutines to finish before main exits.
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Finished goroutines")
}

func sayBaz() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Baz")
}

func IgnoredGoroutineExample() {
	fmt.Println("\nIgnored Goroutine Example:")
	go fmt.Println("Hello")

	// The main function exits before the goroutine can run. So the output will be "World" only.
	fmt.Println("World")
}
