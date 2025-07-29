package main

import (
	"fmt"
	"time"
)

func countTime(start time.Time) time.Duration { return time.Since(start).Round(time.Millisecond) }
func wait(ms time.Duration)                   { time.Sleep(ms * time.Millisecond) }

func selectExample() {
	fmt.Println("\nSelect Example:")
	tickChan, tackChan := make(chan time.Duration), make(chan time.Duration)

	start := time.Now()
	end := time.After(time.Second)

	go func() {
		for range 8 {
			wait(100)
			tickChan <- countTime(start)
		}
		close(tickChan)
	}()

	go func() {
		for range 4 {
			wait(150)
			tackChan <- countTime(start)
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
			fmt.Printf("- [%s]: BOOM!\n", countTime(start))
			return
		default:
			fmt.Printf("- [%s]: ...\n", countTime(start))
			wait(50)
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
			wait(150)
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
