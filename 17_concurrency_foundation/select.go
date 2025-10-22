package main

import (
	"fmt"
	"time"
)

func selectExample() {
	ch1 := make(chan time.Duration)
	ch2 := make(chan time.Duration)

	start := time.Now()
	elapsedFn := func() time.Duration { return time.Since(start).Round(time.Millisecond) }
	end := time.After(time.Second)

	go func() {
		for range 8 {
			time.Sleep(100 * time.Millisecond)
			ch1 <- elapsedFn()
		}
		close(ch1)
	}()

	go func() {
		for range 4 {
			time.Sleep(150 * time.Millisecond)
			ch2 <- elapsedFn()
		}
		close(ch2)
	}()

	for {
		// `select` is like a `switch` statement for channels.
		// It waits for one of the channels to be ready for communication.
		select {
		case value1, ok := <-ch1:
			if ok {
				fmt.Printf("- [%s] ch1: Tick\n", value1)
			} else {
				fmt.Println("- (ch1 closed)")
				ch1 = nil // Avoid select receiving from closed channel
			}
		case value2, ok := <-ch2:
			if ok {
				fmt.Printf("- [%s] ch2: Tack\n", value2)
			} else {
				fmt.Println("- (ch2 closed)")
				ch2 = nil // Same here
			}
		case <-end: // Timeout case
			fmt.Printf("- [%s]: BOOM!\n", elapsedFn())
			return
		default: // No channel is ready
			fmt.Printf("- [%s]: ...\n", elapsedFn())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func sendSelectExample() {
	doneCh := make(chan struct{})
	ch := make(chan int)

	go func() {
		defer close(doneCh)
		for value := range ch {
			fmt.Printf("- Received: %d\n", value)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	go func() {
		defer close(ch)
		for _, value := range []int{10, 20, 30, 40, 50} {
			select {
			case ch <- value: // Send data, if ready to receive
				fmt.Printf("- Sent: %d\n", value)
			case <-time.After(100 * time.Millisecond): // If not ready skip sending
				fmt.Printf("- Skipped %d (Timeout)\n", value)
			}
		}
	}()

	<-doneCh // Wait for receiver to finish
}
