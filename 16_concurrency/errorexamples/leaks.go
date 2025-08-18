package errorexamples

import (
	"context"
	"fmt"
	"time"
)

func leaky() <-chan int {
	ch := make(chan int)
	// Leak: goroutine blocks forever trying to send if nobody receives
	go func() {
		defer fmt.Println("- leaky producer exited (never happens)")
		for i := 0; ; i++ {
			ch <- i // blocks if no receiver -> leak
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return ch
}

func fixed(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
	return out
}

func GoroutineLeakExample() {
	fmt.Println("\nGoroutine Leak Example:")

	_ = leaky() // This starts a leaking goroutine
	time.Sleep(150 * time.Millisecond)
	fmt.Println("- leaked goroutine is still running... (bad)")

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	for v := range fixed(ctx) {
		fmt.Println("- got:", v)
	}

	fmt.Println("- producer stopped cleanly")
}
