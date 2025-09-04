package syncexamples

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randomSleep() {
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
}

func WaitGroupExample() {
	fmt.Println("\nWaitGroup Example:")

	var wg sync.WaitGroup

	// This is a counter that tracks how many goroutines are running.
	// We add to the counter before starting each goroutine and decrement it when the goroutine
	wg.Add(2)

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

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All downloads completed!")
}

func ModernWaitGroupExample() {
	fmt.Println("\nModern WaitGroup Example:")

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

var locations = map[string]string{
	"New York": "USA",
	"Paris":    "France",
	"Tokyo":    "Japan",
	"London":   "UK",
}

func dummyFetch(name string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("- Fetching location for %s...\n", name)
	randomSleep()
	if loc, ok := locations[name]; ok {
		ch <- fmt.Sprintf("%s is in %s", name, loc)
	} else {
		ch <- fmt.Sprintf("Location for %s not found", name)
	}
}

func WaitGroupFetchExample() {
	ch := make(chan string)
	var wg sync.WaitGroup
	cities := []string{"New York", "Paris", "Tokyo", "London", "Berlin"}

	// You could also use wg.Add(1) in the loop but here this is done all at once
	wg.Add(len(cities))

	for _, city := range cities {
		go dummyFetch(city, ch, &wg)
	}

	go func() {
		// Wait in a separate goroutine so main can receive from `ch` immediately.
		wg.Wait()
		close(ch)
	}()

	for loc := range ch {
		fmt.Println("-", loc)
	}
}
