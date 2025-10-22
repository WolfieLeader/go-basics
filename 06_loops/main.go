package main

import (
	"fmt"
	"math"
)

func backwardForExample() {
	fmt.Print("- ")
	for i := 5; i > 0; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func forwardForExample() {
	fmt.Print("- ")
	for i := range 5 { // Modern syntax for `i := 0; i < 5; i++`
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func loopOverEvenNumbers() {
	fmt.Print("- ")
	for i := range 11 {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func whileLoopExample() {
	fmt.Print("- ")
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i *= 2
	}
	fmt.Println()
}

func labelsExample() {
outer:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("  visiting (i=%d, j=%d)\n", i, j)

			// Skip to next outer loop iteration when sum is 2
			if i+j == 2 {
				fmt.Println("  -> continue outer (i+j==2)")
				continue outer
			}

			// Break out of outer loop once we reach (3,0)
			if i == 3 && j == 0 {
				fmt.Println("  -> break outer")
				break outer
			}
		}
	}
}

func gotoExample() {
	target := 9
	for _, n := range []int{2, 4, 6, 8, 9, 10} {
		// Jump to 'found' label when target is found
		if n == target {
			goto found
		}
	}
	fmt.Println("not found")
	return

found:
	fmt.Printf("found %d (jumped with goto)\n", target)
}

func sqrt(x float64) (float64, int) {
	const THRESHOLD = 1e-10
	z, iterations := 1.0, 0

	// Infinite loop until the condition is met
	for {
		iterations++
		prev := z
		z = z - (z*z-x)/(2*z) // Newton's method formula for square root

		if math.Abs(z-prev) < THRESHOLD {
			break
		}
	}
	return z, iterations
}

func main() {
	fmt.Println("Backward For Loop Example:")
	backwardForExample()
	fmt.Println()

	fmt.Println("Forward For Loop Example:")
	forwardForExample()
	fmt.Println()

	fmt.Println("Even For Loop Example:")
	loopOverEvenNumbers()
	fmt.Println()

	fmt.Println("While Loop Example:")
	whileLoopExample()
	fmt.Println()

	fmt.Println("Labels Example:")
	labelsExample()
	fmt.Println()

	fmt.Println("Goto Example:")
	gotoExample()
	fmt.Println()

	fmt.Println("Square root calculation using Newton's method:")
	result, iterations := sqrt(11)
	fmt.Printf("Square root of 11: %.2f (found in %d iterations)\n", result, iterations)
}
