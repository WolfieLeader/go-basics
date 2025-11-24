package basic

import (
	"fmt"
	"time"
)

func SelectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	start := time.Now()
	end := time.After(time.Second)

	go func() {
		defer close(ch1)
		for range 8 {
			time.Sleep(100 * time.Millisecond)
			ch1 <- fmt.Sprintf("- [%s] Tick", time.Since(start).Round(time.Millisecond))
		}
	}()

	go func() {
		defer close(ch2)
		for range 4 {
			time.Sleep(150 * time.Millisecond)
			ch2 <- fmt.Sprintf("- [%s] Tack", time.Since(start).Round(time.Millisecond))
		}
	}()

	for {
		// `select` is like a `switch` statement for channels.
		// It waits for one of the channels to be ready for communication.
		select {
		case v1, ok1 := <-ch1:
			if !ok1 {
				fmt.Println("- (ch1 closed)")
				ch1 = nil // Avoid select receiving from closed channel
				continue
			}
			fmt.Println(v1)
		case v2, ok2 := <-ch2:
			if !ok2 {
				fmt.Println("- (ch2 closed)")
				ch2 = nil // Avoid select receiving from closed channel
				continue
			}
			fmt.Println(v2)
		case <-end: // Timeout case
			fmt.Print("- BOOM!\n")
			return
		default: // No channel is ready, run default case
			fmt.Printf("- [%s]: ...\n", time.Since(start).Round(time.Millisecond))
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func SelectSendExample() {
	ch := make(chan int)
	doneCh := make(chan struct{})

	go func() {
		defer close(doneCh)
		for value := range ch {
			fmt.Printf("- Received: %d\n", value)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	for _, value := range []int{10, 20, 30, 40, 50} {
		select {
		case ch <- value: // Send data, if ready to receive
		case <-time.After(100 * time.Millisecond): // If not ready skip sending
			fmt.Printf("- Skipped %d (Timeout)\n", value)
		}
	}
	close(ch)

	<-doneCh // Wait for receiver to finish
}
