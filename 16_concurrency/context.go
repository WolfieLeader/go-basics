package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sleepBetween(from int, to int) {
	d := time.Duration(from+rand.Intn(to-from+1)) * time.Millisecond
	time.Sleep(d)
}

// Simulate a slow operation that may or may not finish before the context deadline.
// Uses a cancelable timer to avoid leaking a running timer if ctx finishes first.
func slow(ctx context.Context) error {
	timer := time.NewTimer(500 * time.Millisecond)
	defer timer.Stop()

	select {
	case <-timer.C: // Finished work (C is a channel)
		return nil
	case <-ctx.Done(): // The context stopped
		return ctx.Err() // context.Canceled or context.DeadlineExceeded
	}
}

func contextTimeoutExample() {
	fmt.Println("\nContext Timeout Example:")

	// Create a context that cancels automatically after 300ms.
	// Always call cancel: it releases resources even if the deadline has passed.
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	if n := rand.Intn(2); n == 0 {
		cancel() // Simulate an explicit cancellation
	}

	if err := slow(ctx); err != nil {
		switch err {
		case context.DeadlineExceeded:
			fmt.Println("- canceled: deadline exceeded")
		case context.Canceled:
			fmt.Println("- canceled: explicitly canceled")
		default:
			fmt.Println("- canceled:", err)
		}
		return
	}

	fmt.Println("- finished before deadline")
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	count := 1
	for {
		select {
		case <-ctx.Done():
			// Always check ctx.Err() to know why we stopped.
			fmt.Printf("- worker %d: stopping (%v)\n", id, ctx.Err())
			return
		default:
			fmt.Printf("- worker %d: working for the %d time...\n", id, count)
			count++
			sleepBetween(100, 200)
		}
	}
}

func contextCancellationExample() {
	fmt.Println("\nContext Cancellation Example:")

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	for i := range 3 {
		wg.Add(1)
		go worker(ctx, i+1, &wg)
	}

	// Cancel after a random time
	sleepBetween(200, 600)
	cancel() // Broadcast stop to all workers

	// Wait deterministically for all workers to observe ctx.Done() and exit.
	wg.Wait()
	fmt.Println("All workers canceled")
}

// Best practice: use a private, typed key to avoid collisions with other packages.
type ctxKey string

const requestIDKey ctxKey = "requestID"

func reqLog(ctx context.Context, msg string) {
	if v := ctx.Value(requestIDKey); v != nil {
		fmt.Printf("[id:%s] %s\n", v, msg)
		return
	}
	fmt.Println(msg)
}

func doWork(ctx context.Context) {
	reqLog(ctx, "start")
	time.Sleep(50 * time.Millisecond)
	reqLog(ctx, "done")
}

func contextValuesExample() {
	fmt.Println("\nContext Values Example:")

	// Attach a request ID to the context.
	// Note: keep context values small and immutable; prefer passing explicit params for large data.
	ctx := context.WithValue(context.Background(), requestIDKey, "req-12345")

	doWork(ctx)
}
