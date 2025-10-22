package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func slowOperation(ctx context.Context) error {
	timer := time.NewTimer(250 * time.Millisecond)
	defer timer.Stop()

	select {
	case <-timer.C: // Finished work (C is a channel)
		return nil
	case <-ctx.Done(): // The context stopped
		return ctx.Err() // `context.Canceled` or `context.DeadlineExceeded`
	}
}

func contextTimeoutExample() {
	duration := time.Duration(200+rand.Intn(300)) * time.Millisecond
	fmt.Printf("- Timeout set to %v\n", duration)

	// Create a context that cancels automatically after 200ms to 500ms
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel() // Call cancel to release resources

	if randomNumber := rand.Intn(3); randomNumber == 0 {
		fmt.Println("- Simulating explicit cancellation")
		cancel() // Simulate an explicit cancellation
	}

	if err := slowOperation(ctx); err != nil {
		switch err {
		case context.DeadlineExceeded:
			fmt.Println("- Deadline exceeded")
		case context.Canceled:
			fmt.Println("- Explicit cancellation")
		default:
			fmt.Printf("- Canceled: %v\n", err)
		}
		return
	}

	fmt.Println("- Finished before deadline!")
}
