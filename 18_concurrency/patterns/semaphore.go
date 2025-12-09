package patterns

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func SemaphoreExample() {
	const maxWorkers = 5
	results := make(chan float64)
	var wg sync.WaitGroup

	// Semaphore channel to limit concurrency
	sem := make(chan struct{}, maxWorkers)

	pows := generator(1, 2, 4, 8, 16, 32, 64, 128, 256, 512)

	// Fan-Out
	for num := range pows {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sem <- struct{}{} // Acquire semaphore

			time.Sleep(50 * time.Millisecond)
			results <- math.Log2(float64(num))

			<-sem // Release semaphore
		}(num)
	}

	// Wait
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-In
	nums := make([]float64, 0)
	for r := range results {
		nums = append(nums, r)
	}
	fmt.Printf("- Numbers (limited to 5 gorotinues): %v\n", nums)
}
