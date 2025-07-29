package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, limit int) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := range limit {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(c) // Close the channel when done to signal completion.
	}()
	return c // Return the channel to the caller.
}

func main() {
	// goroutineExample()
	// ignoredGoroutineExample()
	// channelExample()
	// channelIterationExample()
	// channelDirectionExample()
	selectExample()
}
