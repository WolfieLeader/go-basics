package main

import "fmt"

func main() {
	num, err := convertToInt("42")
	if err != nil {
		fmt.Println("Error converting to int:", err)
		return
	}
	fmt.Printf("Converted string to int: %d\n", num)
	printOS()
	loopBackward(5)
	loopForward(5)
	printEvenNumbers(23)
	whileLoop(5)
	fmt.Printf("Square root of 11: %.2f\n", sqrt(11))
}
