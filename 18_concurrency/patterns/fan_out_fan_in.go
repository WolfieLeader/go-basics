package patterns

import (
	"fmt"
	"sync"
	"time"
)

func FanOutFanInExample() {
	results := make(chan int)
	var wg sync.WaitGroup

	// Spread out work
	tens := generator(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)

	// Fan-Out - Start workers to split work across multiple goroutines
	for num := range tens {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond) // Simulate work
			results <- num / 10
		}(num)
	}

	// Wait for workers to finish and close channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-In - Collect all computed results
	nums := make([]int, 0)
	for r := range results {
		nums = append(nums, r)
	}
	fmt.Printf("- Numbers (10 gorotinues as the number of tasks): %v\n", nums)
}
