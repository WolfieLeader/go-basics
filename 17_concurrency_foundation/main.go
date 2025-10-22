package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepFromTo(from int, to int) {
	time.Sleep(time.Duration(from+rand.Intn(to-from+1)) * time.Millisecond)
}

func main() {
	fmt.Println("Goroutine Example:")
	goroutineExample()
	fmt.Println()

	fmt.Println("Goroutine Without Wait Example:")
	goroutineWithoutWaitExample()
	fmt.Println()

	fmt.Println("Unbuffered Channel Example:")
	unbufferedChannelExample()
	fmt.Println()

	fmt.Println("Buffered Channel Example:")
	bufferedChannelExample()
	fmt.Println()

	fmt.Println("Comma-Ok Channel Example:")
	commaOkChannelExample()
	fmt.Println()

	fmt.Println("Channel Directions Example:")
	channelDirectionsExample()
	fmt.Println()

	fmt.Println("Select Example:")
	selectExample()
	fmt.Println()

	fmt.Println("Send Select Example:")
	sendSelectExample()
	fmt.Println()

	fmt.Println("Sync Wait Group Example:")
	syncWaitGroupExample()
	fmt.Println()

	fmt.Println("Modern Sync Wait Group Example:")
	modernSyncWaitGroupExample()
	fmt.Println()

	fmt.Println("Sync Mutex Example:")
	syncMutexExample()
	fmt.Println()

	fmt.Println("Context Timeout Example:")
	contextTimeoutExample()
	fmt.Println()

	fmt.Println("Context Cancellation Example:")
	contextCancellationExample()
	fmt.Println()

	fmt.Println("Context Value Example:")
	contextValueExample()
	fmt.Println()
}
