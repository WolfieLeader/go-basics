package main

import (
	"example/greetings" // go mod edit -replace example/greetings=../greetings
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
