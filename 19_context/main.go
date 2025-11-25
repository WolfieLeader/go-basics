package main

import "fmt"

func main() {
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