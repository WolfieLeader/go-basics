package patterns

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func FanOutFanInExample() {
	const workers = 3
	nums := []float64{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100, 121, 144, 169, 196, 225, 256, 289, 324, 361, 400}

	jobs := make(chan float64)
	results := make(chan float64)

	var wg sync.WaitGroup

	// Fan-Out - Start workers to split work across multiple goroutines
	wg.Add(workers)
	for w := 1; w <= workers; w++ {
		go worker(&wg, jobs, results)
	}

	// Spread the work
	go func() {
		for _, n := range nums {
			jobs <- n
		}
		close(jobs)
	}()

	// Wait for workers to finish and close channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-In - Collect all computed results
	roots := make([]float64, 0, len(nums))
	for r := range results {
		roots = append(roots, r)
	}
	fmt.Printf("- %v", roots)
}

func worker(wg *sync.WaitGroup, jobs <-chan float64, results chan<- float64) {
	defer wg.Done()
	for num := range jobs {
		time.Sleep(50 * time.Millisecond) // Simulate work
		results <- math.Sqrt(num)
	}
}
