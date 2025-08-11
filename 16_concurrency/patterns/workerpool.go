package patterns_examples

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	id int
}

type result struct {
	id int
	ok bool
}

func worker(id int, jobs <-chan job, results chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(60 * time.Millisecond)
		results <- result{j.id, true}
		fmt.Printf("- worker %d finished job %d\n", id, j.id)
	}
}

func WorkerPoolExample() {
	fmt.Println("\nWorker Pool Example:")

	jobs := make(chan job)
	results := make(chan result)

	var wg sync.WaitGroup
	const totalWorkers = 3
	const totalJobs = 8

	for w := range totalWorkers {
		wg.Add(1)
		go worker(w+1, jobs, results, &wg)
	}

	// Feed jobs
	go func() {
		for j := range totalJobs {
			jobs <- job{j + 1}
		}
		close(jobs)
	}()

	// Close results after workers exit
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		_ = r // consume results
	}
}
