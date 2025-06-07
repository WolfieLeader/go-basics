package main

import "fmt"

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println(RecursiveFibonacci(4))
}

//TODO: https://go.dev/tour/methods/1
