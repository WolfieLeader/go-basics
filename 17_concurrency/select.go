package main

import (
	"fmt"
	"time"
)

func selectExample() {
	fmt.Println("\nSelect Example:")
	tickChan, tackChan := make(chan time.Duration), make(chan time.Duration)

	start := time.Now()
	countTime := func() time.Duration { return time.Since(start).Round(time.Millisecond) }
	end := time.After(time.Second)

	go func() {
		for range 8 {
			time.Sleep(100 * time.Millisecond)
			tickChan <- countTime()
		}
		close(tickChan)
	}()

	go func() {
		for range 4 {
			time.Sleep(150 * time.Millisecond)
			tackChan <- countTime()
		}
		close(tackChan)
	}()

	for {
		// `select` is like a `switch` statement for channels.
		// It waits for one of the channels to be ready for communication.
		select {
		case v1, ok := <-tickChan:
			if ok {
				fmt.Printf("- [%s] ch1: Tick\n", v1)
			} else {
				fmt.Println("- (ch1 closed)")
				// Avoid select receiving from closed channel
				tickChan = nil
			}
		case v2, ok := <-tackChan:
			if ok {
				fmt.Printf("- [%s] ch2: Tack\n", v2)
			} else {
				fmt.Println("- (ch2 closed)")
				tackChan = nil
			}
		case <-end:
			fmt.Printf("- [%s]: BOOM!\n", countTime())
			return
		default:
			fmt.Printf("- [%s]: ...\n", countTime())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func selectSendExample() {
	fmt.Println("\nSelect Sending Example:")

	// Using a done channel to signal when all messages have been processed.
	done := make(chan struct{})
	data := make(chan byte)

	go func() {
		for msg := range data {
			fmt.Printf("- Received: %c\n", msg)
			time.Sleep(150 * time.Millisecond)
		}
		close(done)
	}()

	go func() {
		for _, letter := range []byte{'A', 'B', 'C', 'D', 'E'} {
			select {
			// Will send to the data channel if it's ready to receive
			case data <- letter:
				fmt.Printf("- Sent: %c\n", letter)
			// If 100 milliseconds pass without the channel being ready, skip sending
			case <-time.After(100 * time.Millisecond):
				fmt.Printf("- Skipped %c (Timeout)\n", letter)
			}
		}
		close(data)
	}()

	// Wait for the goroutine to finish processing all messages
	<-done
}
