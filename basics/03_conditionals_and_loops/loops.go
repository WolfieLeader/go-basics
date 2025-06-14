package main

import "fmt"

func loopBackward(n int) {
	fmt.Println("Looping backward:")
	// initialization statement, condition, and post statement are all optional in Go's for loop
	for j := n; j > 0; j-- {
		fmt.Printf("- j = %d\n", j)
	}
}

func loopForward(n int) {
	fmt.Println("Looping forward:")
	for i := range n {
		fmt.Printf("- i = %d\n", i)
	}
}

func printEvenNumbers(n int) {
	fmt.Println("Looping over even numbers:")
	for k := range n {
		if k%2 != 0 {
			continue // Skip odd numbers
		}
		fmt.Printf("- k = %d\n", k)
	}
}

func whileLoop(n int) {
	fmt.Println("While loop:")
	l := 0
	// In Go, the for loop can also be used as a while loop
	for l < n {
		fmt.Printf("- l = %d\n", l)
		l++
	}
}

// This is how you can implement a square root function using Newton's method
func sqrt(x float64) float64 {
	const THRESHOLD = 1e-10
	z, counter := 1.0, 0

	// Infinite loop until the condition is met
	for {
		counter++
		prev := z
		z = z - (z*z-x)/(2*z) // Newton's method formula for square root

		if absolute(z-prev) < THRESHOLD {
			break // Exit the loop if the change is within the threshold
		}
	}
	fmt.Printf("Sqrt converged after %d iterations\n", counter)
	return z
}
