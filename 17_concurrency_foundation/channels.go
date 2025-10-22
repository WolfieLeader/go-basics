package main

import (
	"fmt"
	"time"
)

// Channels are used to communicate between goroutines.
// Unbuffered channel - blocks until both sender and receiver are ready.
func unbufferedChannelExample() {
	ch := make(chan int)
	arr := []int{1, 2, 3, 4, 5}

	go func() {
		for _, value := range arr {
			ch <- value // Send value to unbuffered channel.
			fmt.Printf("- Sent: %d (len=%d)\n", value, len(ch))
		}
		close(ch) // Signal that no more values will be sent.
	}()

	for range len(arr) {
		received := <-ch // Block until value is received.
		fmt.Printf("- Received: %d\n", received)
	}

}

// Buffered channel - does not block until the buffer is full.
func bufferedChannelExample() {
	ch := make(chan byte, 2) // Buffered channel with capacity of 2.

	for _, value := range []byte{'A', 'B'} {
		ch <- value // The channel can hold 2 values without blocking.
		fmt.Printf("- Sent: %q (len=%d)\n", value, len(ch))
	}

	go func() {
		for _, value := range []byte{'C', 'D', 'E', 'F', 'G'} {
			ch <- value
			fmt.Printf("- Sent: %q (len=%d)\n", value, len(ch))
		}
		close(ch)
	}()

	// Receive values from the buffered channel.
	for value := range ch {
		fmt.Printf("- Received: %q\n", value)
	}
}

func commaOkChannelExample() {
	ch := make(chan int, 1)
	ch <- 50
	close(ch)

	// Check if channel is closed using the comma-ok idiom.
	value, ok := <-ch
	fmt.Printf("- Comma-ok receive: %d (ok: %t)\n", value, ok)

	value, ok = <-ch
	fmt.Printf("- Comma-ok receive (after close): %d (ok: %t)\n", value, ok)
}

func readOnlyFn[T comparable](ch <-chan T) {
	fmt.Printf("- Read from channel (read-only): %v\n", <-ch)
}

func writeOnlyFn[T comparable](ch chan<- T, value T) {
	fmt.Printf("- Wrote to channel (write-only): %v\n", value)
	ch <- value
}

func readWriteFn[T comparable](ch chan T, values ...T) {
	if len(values) == 0 {
		fmt.Printf("- Read from channel (both): %v\n", <-ch)
	} else {
		fmt.Printf("- Wrote to channel (both): %v\n", values[0])
		ch <- values[0]
	}
}

func channelDirectionsExample() {
	ch := make(chan int)

	go writeOnlyFn(ch, 123)
	go readOnlyFn(ch)
	time.Sleep(100 * time.Millisecond)

	go readWriteFn(ch, 456)
	go readWriteFn(ch)
	time.Sleep(100 * time.Millisecond)
}
