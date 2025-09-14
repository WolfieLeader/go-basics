package main

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/14_packages/utils"
)

func main() {
	str := "Hello World! 🚀"
	// Using the Reverse function from the utils package
	fmt.Printf("- %s -> %s\n", str, utils.Reverse(str))
	// Since privateFunction is not exported (starts with a lowercase letter), it cannot be accessed here.
}

// Init function executes before main
func init() {
	fmt.Println("Init function in main.go 🪙")
}
