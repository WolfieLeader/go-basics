package patterns

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func FanOutFanInExample() {
	const workers = 3
	results := make(chan int)
	var wg sync.WaitGroup

	// Spread out work (via channel) with generator pattern to produce perfect squares
	tasks := squareGenerator(20)

	// Fan-Out - Start workers to split work across multiple goroutines
	wg.Add(workers)
	for w := 1; w <= workers; w++ {
		go sqrtWorker(&wg, tasks, results)
	}

	// Wait for workers to finish and close channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-In - Collect all computed results
	roots := make([]int, 0)
	for r := range results {
		roots = append(roots, r)
	}
	fmt.Printf("- %v", roots)
}

func squareGenerator(max int) <-chan float64 {
	ch := make(chan float64)
	go func() {
		defer close(ch)
		for i := 0; i <= max; i++ {
			ch <- float64(i * i)
		}
	}()
	return ch
}

func sqrtWorker(wg *sync.WaitGroup, tasks <-chan float64, results chan<- int) {
	defer wg.Done()
	for num := range tasks {
		time.Sleep(50 * time.Millisecond) // Simulate work
		results <- int(math.Sqrt(num))
	}
}
