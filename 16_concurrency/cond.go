package main

import (
	"fmt"
	"sync"
	"time"
)

func condExample() {
	lock := sync.Mutex{}
	// `sync.Cond` is used to wait for and signal conditions between goroutines.
	// It requires a mutex to protect shared state.
	cond := sync.NewCond(&lock)

	queue := []string{}
	const capacity = 3

	// Consumer goroutine
	go func() {
		for {
			cond.L.Lock() // Lock before accessing queue

			// If the queue is empty, wait for a signal from producer
			for len(queue) == 0 {
				cond.Wait() // Unlocks and waits for signal, re-locks upon wake
			}

			// Remove item from front of queue
			item := queue[0]
			queue = queue[1:]

			fmt.Println("  Consumed:", item)

			// Notify producers there may be room now
			cond.Signal()
			cond.L.Unlock()

			time.Sleep(300 * time.Millisecond)
		}
	}()

	// Producer goroutine
	for i := range 5 {
		cond.L.Lock() // Lock before accessing queue

		// If the queue is full, wait for consumer to remove something
		for len(queue) >= capacity {
			cond.Wait()
		}

		// Add item to queue
		item := fmt.Sprintf("item-%d", i+1)
		queue = append(queue, item)

		fmt.Println("- Produced:", item)

		// Notify consumers there's something to consume
		cond.Signal()
		cond.L.Unlock()

		time.Sleep(200 * time.Millisecond)
	}
}
