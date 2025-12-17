package basic

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func SyncWaitGroupExample() {
	var wg sync.WaitGroup
	downloadFn := func(item string) {
		fmt.Printf("- Downloading %s...\n", item)
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Printf("- %s downloaded!\n", item)
	}

	// `wg.Add` increments the WaitGroup counter by the number of goroutines to wait for
	wg.Add(2)

	// `wg.Done` decrements the counter when the goroutine completes
	go func() { defer wg.Done(); downloadFn("Movie") }()
	go func() { defer wg.Done(); downloadFn("Music") }()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("- All downloads completed!")
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
