package patterns

import (
	"fmt"
	"sync"
	"time"
)

func doWork(id int, sem <-chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		<-sem     // release
		wg.Done() // signal that the goroutine is done
	}()

	fmt.Printf("- task %d\n", id)
	time.Sleep(750 * time.Millisecond)
}

// Semaphore pattern is a concurrency pattern that restricts the number of goroutines
// that can access a resource or perform a task simultaneously.
func SemaphoreExample() {
	fmt.Println("\nSemaphore Example:")

	var wg sync.WaitGroup
	const maxConcurrent = 3
	const totalTasks = 15
	sem := make(chan struct{}, maxConcurrent) // capacity = permits

	for i := range totalTasks {
		task := i + 1
		wg.Add(1)

		sem <- struct{}{} // acquire
		go doWork(task, sem, &wg)
	}

	wg.Wait()
}
