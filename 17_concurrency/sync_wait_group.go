package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func syncWaitGroupExample() {
	var wg sync.WaitGroup

	wg.Add(2) // We have 2 goroutines to wait for

	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		fmt.Println("- Downloading Movie...")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("- Movie downloaded!")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("- Downloading Music...")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("- Music downloaded!")
	}()

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("- All downloads completed!")
}

func modernSyncWaitGroupExample() {
	var wg sync.WaitGroup

	for id := 1; id <= 3; id++ {
		// `wg.Go` is a convenience method that adds 1 to the WaitGroup counter,
		// starts a new goroutine, and decrements the counter when the goroutine finishes.
		wg.Go(func() {
			fmt.Printf("- Worker %d started\n", id)
			time.Sleep(time.Duration(id) * 500 * time.Millisecond)
			fmt.Printf("- Worker %d done\n", id)
		})
	}

	wg.Wait()
	fmt.Println("- All workers completed!")
}
