package patterns_examples

import (
	"fmt"
	"time"
)

func doWork(id int, sem chan struct{}) {
	sem <- struct{}{} // acquire

	defer func() {
		<-sem
	}() // release

	fmt.Printf("- task %d started\n", id)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("- task %d done\n", id)
}

// SemaphoreExample demonstrates the use of a semaphore to limit concurrency.
func SemaphoreExample() {
	fmt.Println("\nSemaphore Example:")

	const maxConcurrent = 2
	const totalTasks = 5
	sem := make(chan struct{}, maxConcurrent) // capacity = permits

	for i := range totalTasks {
		go doWork(i, sem)
	}

	time.Sleep(time.Second)
}
