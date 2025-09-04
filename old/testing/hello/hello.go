package main

import (
	"fmt"
	"log"

	// go mod edit -replace example/greetings=../greetings
	// go mod tidy
	"example/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
