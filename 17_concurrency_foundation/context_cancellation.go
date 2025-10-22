package main

import (
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	count := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("- worker %d: stopping (ctx done): %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("- worker %d: working for the %d time...\n", id, count)
			count++
			sleepFromTo(100, 200)
		}
	}
}

const WORKERS = 3

func contextCancellationExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(WORKERS)
	for id := 1; id <= WORKERS; id++ {
		go worker(ctx, id, &wg)
	}

	sleepFromTo(200, 600)
	cancel() // Cancel the context to stop all workers

	wg.Wait()
	fmt.Println("- All workers stopped")
}
