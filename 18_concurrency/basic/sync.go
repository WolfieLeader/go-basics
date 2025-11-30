package basic

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func SyncWaitGroupExample() {
	var wg sync.WaitGroup

	// `wg.Add` increments the WaitGroup counter by the number of goroutines to wait for
	wg.Add(2)

	// `wg.Done` decrements the counter when the goroutine completes
	go func() { defer wg.Done(); downloadFn("Movie") }()
	go func() { defer wg.Done(); downloadFn("Music") }()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("- All downloads completed!")
}

func downloadFn(x string) {
	fmt.Printf("- Downloading %s...\n", x)
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Printf("- %s downloaded!\n", x)
}

func ModernSyncWaitGroupExample() {
	var wg sync.WaitGroup

	for id := 1; id <= 3; id++ {
		// `wg.Go` is a convenience method that adds 1 to the WaitGroup counter,
		// starts a new goroutine, and decrements the counter when the goroutine finishes.
		wg.Go(func() {
			fmt.Printf("- [worker %d]: Processsing...\n", id)
			t := time.Now()
			time.Sleep(time.Duration(id) * 300 * time.Millisecond)
			fmt.Printf("- [worker %d]: Done! (took %v)\n", id, time.Since(t).Round(time.Millisecond))
		})
	}

	wg.Wait()
	fmt.Println("- All workers completed!")
}

func SyncMutexExample() {
	// Mutex to synchronize access to the counter
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter, racyCounter := 0, 0

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
