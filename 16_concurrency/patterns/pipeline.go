package patterns

import (
	"fmt"
)

// Stage 1: generate numbers
func gen(n ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range n {
			out <- v
		}
	}()
	return out
}

// Stage 2: square
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// Stage 3: add one
func addOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n + 1
		}
	}()
	return out
}

// Pipeline pattern takes a series of stages and connects them
func PipelineExample() {
	fmt.Println("\nPipeline Example:")
	for v := range addOne(square(gen(1, 2, 3, 4))) {
		fmt.Println("-", v)
	}
}
