package basic

import (
	"fmt"
	"sync"
)

func SyncMutexExample() {
	var (
		mu          sync.Mutex // Mutex to synchronize access to the counter
		wg          sync.WaitGroup
		counter     = 0
		racyCounter = 0
	)

	for range 1000 {
		wg.Go(func() {
			mu.Lock() // Lock the mutex before accessing the counter
			counter++
			mu.Unlock() // Unlock the mutex, can also use `defer` for safety
		})

		wg.Go(func() {
			// Not protected, leading to race conditions
			racyCounter++
		})
	}
	wg.Wait()

	fmt.Printf("- Counter with mutex:         %d\n", counter)
	fmt.Printf("- Racy counter without mutex: %d\n", racyCounter)
}
