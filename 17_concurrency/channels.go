package main

import (
	"fmt"
	"time"
)

func channelExample() {
	fmt.Println("\nChannel Example:")
	// To create a channel, use the `make` function with `chan` keyword.

	// This is an unbuffered channel of type `int`.
	// Unbuffered means that a send and a receive operation must happen simultaneously.
	ch1 := make(chan int)

	// This is a buffered channel of type `byte` with a capacity of 2.
	// Buffered channels allow sending values without blocking until the buffer is full.
	ch2 := make(chan byte, 2)

	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := [5]byte{'A', 'B', 'C', 'D', 'E'}

	go func() {
		for _, v := range a1 {
			// Send value to the channel.
			ch1 <- v
		}
	}()

	go func() {
		for _, v := range a2 {
			// Send value to the buffered channel, which can hold up to 2 values without blocking.
			// Once full, further sends will block until a receiver reads a value.
			ch2 <- v
		}
	}()

	// Print the capacities of the channels. Since `ch1` is unbuffered, its capacity is 0.
	fmt.Printf("- ch1 (unbuffered) capacity: %d\n", cap(ch1))
	fmt.Printf("- ch2 (buffered) capacity: %d\n\n", cap(ch2))

	for range 5 {
		// Receiving from channels. This blocks until a value is available in the channel.
		v1 := <-ch1
		v2 := <-ch2

		// `len(channel)` gives the number of values currently buffered (queued) in the channel.
		// For unbuffered channels, this will always be 0.
		fmt.Printf("- ch1: %d (len:%d), ch2: %c (len:%d)\n", v1, len(ch1), v2, len(ch2))
	}

	// Closing the channels to signal that no more values will be sent.
	close(ch1)
	close(ch2)

	ch3 := make(chan string, 1)
	ch3 <- "test"
	close(ch3)

	// Comma-ok idiom to check if the channel is closed and nothing is left to read.
	v, ok := <-ch3
	fmt.Printf("\n- ch3: %s (ok: %t)\n", v, ok)
	v, ok = <-ch3
	fmt.Printf("- ch3: %s (ok: %t)\n", v, ok)
}

func channelIterationExample() {
	fmt.Println("\nChannel Iteration Example:")
	ch := make(chan int)

	go func() {
		for i := range 3 {
			ch <- i
		}
		close(ch) // Close the channel to signal that no more values will be sent.
	}()

	for i := range ch {
		fmt.Printf("- %d\n", i)
	}
}

func channelDirectionExample() {
	fmt.Println("\nChannel Direction Example:")
	ch1, ch2 := make(chan int), make(chan int)

	go writeOnly(ch1, 12)
	go readOnly(ch1)
	go bidirectional(ch2, 34)
	go func() { ch2 <- <-ch2 }() // Echoing the value since the function waits for a value to be sent.

	time.Sleep(100 * time.Millisecond) // Sleep to ensure all goroutines finish before main exits.
	fmt.Println("Finished channel direction example")
}

func readOnly(ch <-chan int) {
	fmt.Printf("- Read-only: %d\n", <-ch)
}

func writeOnly(ch chan<- int, num int) {
	fmt.Printf("- Write-only: Sending %d\n", num)
	ch <- num
}

func bidirectional(ch chan int, num int) {
	fmt.Printf("- Bidirectional: Sending %d\n", num)
	ch <- num

	// Wait for a value to be received from the channel.
	fmt.Printf("- Bidirectional: Received %d\n", <-ch)
}
