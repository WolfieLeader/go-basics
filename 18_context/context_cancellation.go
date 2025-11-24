package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
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
			time.Sleep(time.Duration(100+rand.Intn(100)) * time.Millisecond)
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
	
	time.Sleep(time.Duration(200+rand.Intn(400)) * time.Millisecond)
	cancel() // Cancel the context to stop all workers

	wg.Wait()
	fmt.Println("- All workers stopped")
}
