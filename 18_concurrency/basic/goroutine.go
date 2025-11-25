package basic

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// They are used to perform concurrent (asynchronous) tasks in Go.
func GoroutineExample() {
	// `go` keyword is used to start a new goroutine (created 2 goroutines here).
	go fmt.Println("- Foo")
	go func() { fmt.Println("- Bar") }()

	time.Sleep(100 * time.Millisecond) // Allow goroutines to finish before main exits.
	fmt.Println("Main goroutine finished after waiting.")

	// Since we don't wait for the goroutines to finish, they may not execute.
	go fmt.Println("- Unexecuted Message")
	fmt.Println("Main goroutine finished without waiting.")
}

// Channels are used to communicate between goroutines.

// Unbuffered channel - blocks until both sender and receiver are ready.
func UnbufferedChannelExample() {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i // Send value to channel.
			fmt.Printf("- Sent: %d (len=%d)\n", i, len(ch))
		}
		close(ch) // Signal that no more values will be sent.
	}()

	for range 5 {
		received := <-ch // Block until value is received.
		fmt.Printf("- Received: %d\n", received)
	}

}

// Buffered channel - does not block until the buffer is full.
func BufferedChannelExample() {
	ch := make(chan byte, 2) // Buffered channel with capacity of 2.

	ch <- 'A'
	ch <- 'B'
	fmt.Printf("- Buffered sends complete (len=%d)\n", len(ch))

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

func CommaOkChannelExample() {
	ch := make(chan int, 1)

	ch <- 50
	close(ch)

	v, ok := <-ch
	fmt.Printf("- First receive: %d (ok=%t)\n", v, ok)

	v, ok = <-ch
	fmt.Printf("- Second receive (after close): %d (ok=%t)\n", v, ok)
}

func ChannelDirectionsExample() {
	ch := make(chan int)

	go writeOnlyFn(ch, 123)
	go readOnlyFn(ch)
	time.Sleep(50 * time.Millisecond)

	go readWriteFn(ch, 456)
	go readWriteFn(ch)
	time.Sleep(50 * time.Millisecond)
}

func readOnlyFn(ch <-chan int) {
	value := <-ch
	fmt.Printf("- Read from channel (read-only): %v\n", value)
}

func writeOnlyFn(ch chan<- int, value int) {
	fmt.Printf("- Wrote to channel (write-only): %v\n", value)
	ch <- value
}

func readWriteFn(ch chan int, values ...int) {
	if len(values) == 0 {
		value := <-ch
		fmt.Printf("- Read from channel (read-write): %v\n", value)
		return
	}

	fmt.Printf("- Wrote to channel (read-write): %v\n", values[0])
	ch <- values[0]
}
