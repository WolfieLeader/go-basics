package patterns_examples

import (
	"fmt"
	"sync"
	"time"
)

// Fan-Out / Fan-In splits work across workers (fan-out) and merges outputs (fan-in).
func FanOutFanInExample() {
	fmt.Println("\nFan-out / Fan-in Example:")

	const totalWorkers = 3
	const totalJobs = 6

	var wg sync.WaitGroup
	jobs := make(chan byte)
	results := make(chan string, totalJobs) // buffer to reduce backpressure on workers

	// Fan-out: start workers (consumers)
	for worker := range totalWorkers {
		w := worker + 1 // capture loop variable for the goroutine
		wg.Go(func() {
			for j := range jobs {
				// Simulate work
				time.Sleep(50 * time.Millisecond)
				results <- fmt.Sprintf("worker %d -> job %c", w, j)
			}
		})
	}

	// Producer
	go func() {
		for i := range totalJobs {
			jobs <- byte('A' + i)
		}
		close(jobs) // signal no more jobs
	}()

	// Fan-in: close results when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("-", r)
	}
}
