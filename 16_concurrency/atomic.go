package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	workerCount    = 3
	tasksPerWorker = 1000
)

func atomicExample() {
	var totalTasks int64 // Shared atomic counter
	var wg sync.WaitGroup
	start := time.Now()

	// Launch multiple worker goroutines
	for i := range workerCount {
		workerId := i + 1

		wg.Go(func() {
			for j := range tasksPerWorker {
				// Atomic operation to safely increment the shared counter
				// This ensures that even if multiple goroutines try to update it at the same time
				atomic.AddInt64(&totalTasks, 1)

				// Simulate doing some work
				if (j+1)%500 == 0 {
					fmt.Printf("🔧 Worker %d processed task %d\n", workerId, j+1)
				}
				time.Sleep(time.Microsecond * 500) // Simulate light work
			}
		})
	}

	// Wait for all workers to finish
	wg.Wait()

	// Safely read the final total
	final := atomic.LoadInt64(&totalTasks)
	fmt.Printf("✅ All workers done. Total tasks: %d (expected: %d)\n", final, workerCount*tasksPerWorker)
	fmt.Printf("⏱️ Elapsed: %s\n", time.Since(start).Round(time.Millisecond))
}
