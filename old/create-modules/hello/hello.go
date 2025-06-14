package main

import (
	"fmt"
	// go mod edit -replace example/greetings=../greetings
	// go mod tidy
	"example/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
