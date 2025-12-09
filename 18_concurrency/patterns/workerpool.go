package patterns

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func WorkerPoolExample() {
	const workerCount = 3
	results := make(chan float64)
	var wg sync.WaitGroup

	squares := generator(0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100)

	// Fan-Out - Start FIXED number of workers to split work across multiple goroutines
	for w := 1; w <= workerCount; w++ {
		wg.Go(func() {
			for num := range squares {
				time.Sleep(50 * time.Millisecond)
				results <- math.Sqrt(float64(num))
			}
		})
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
	fmt.Printf("- Numbers (3 gorotinues): %v\n", nums)
}
