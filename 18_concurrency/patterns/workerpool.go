package patterns

import (
	"fmt"
	"sync"
	"time"
)

func WorkerPoolExample() {
	const taskCount = 10
	const workerCount = 3

	results := make(chan int)
	var wg sync.WaitGroup

	// Spread out work (via channel) with generator pattern to produce doubled numbers
	tens := tenXGenerator(taskCount)

	// Fan-Out - Start FIXED number of workers to split work across multiple goroutines
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for t := range tens {
				fmt.Printf("- [worker %d]: Processing %d\n", workerId, t)
				time.Sleep(50 * time.Millisecond) // Simulate work
				results <- t / 10
			}
		}(w)
	}

	// Wait for workers to finish and close channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-In - Collect all computed results
	nums := make([]int, 0)
	for n := range results {
		nums = append(nums, n)
	}
	fmt.Printf("- Numbers: %v\n", nums)
}

func tenXGenerator(count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= count; i++ {
			ch <- i * 10
		}
	}()
	return ch
}
