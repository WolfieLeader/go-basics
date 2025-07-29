package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomSleep() {
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
}

func waitGroupExample() {
	fmt.Println("\nWaitGroup Example:")

	var wg sync.WaitGroup

	// This is a counter that tracks how many goroutines are running.
	// We add to the counter before starting each goroutine and decrement it when the goroutine
	wg.Add(3)

	go func() {
		// At the end of the goroutine, we call Done to decrement the counter
		defer wg.Done()
		fmt.Println("- Downloading Movie...")
		randomSleep()
		fmt.Println("- Movie downloaded!")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("- Downloading Music...")
		randomSleep()
		fmt.Println("- Music downloaded!")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("- Downloading Game...")
		randomSleep()
		fmt.Println("- Game downloaded!")
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All downloads completed!")
}

func newWaitGroupExample() {
	fmt.Println("\nNew WaitGroup Example:")

	var wg sync.WaitGroup

	for i := range 3 {
		id := i + 1
		// `wg.Go` is a convenience method that adds 1 to the WaitGroup counter,
		// starts a new goroutine, and decrements the counter when the goroutine finishes.
		wg.Go(func() {
			fmt.Printf("- Worker %d started\n", id)
			time.Sleep(time.Duration(id) * 500 * time.Millisecond)
			fmt.Printf("- Worker %d done\n", id)
		})
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All workers completed!")
}
