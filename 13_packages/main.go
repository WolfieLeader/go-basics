package main

import "github.com/WolfieLeader/go-basics/13_packages/utils"

func main() {
	// Using the Reverse function from the utils package
	reversed := utils.Reverse("Hello, World🔥⚡🚀")
	println("Reversed string:", reversed)
	// Since privateFunction is not exported (starts with a lowercase letter), it cannot be accessed here.
}
