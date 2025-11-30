package patterns

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func FanOutFanInExample() {
	const taskCount = 20
	results := make(chan float64)
	var wg sync.WaitGroup

	// Spread out work (via channel) with generator pattern to produce perfect squares
	squares := squareGenerator(taskCount)

	// Fan-Out - Start workers to split work across multiple goroutines
	for sq := range squares {
		wg.Add(1)
		go func(num float64) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond) // Simulate work
			results <- math.Sqrt(num)
		}(sq)
	}

	// Wait for workers to finish and close channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-In - Collect all computed results
	roots := make([]float64, 0)
	for r := range results {
		roots = append(roots, r)
	}
	fmt.Printf("- Roots of numbers: %v\n", roots)
}

func squareGenerator(max int) <-chan float64 {
	ch := make(chan float64)
	go func() {
		defer close(ch)
		for i := 1; i <= max; i++ {
			ch <- float64(i * i)
		}
	}()
	return ch
}
