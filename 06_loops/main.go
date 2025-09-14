package main

import "fmt"

func backwardForExample() {
	fmt.Println("\nBackward For Loop Example:")
	// Looping backward from 5 to 1
	fmt.Print("- ")
	for i := 5; i > 0; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func forwardForExample() {
	fmt.Println("\nForward Range For Loop Example:")
	// Looping forward from 1 to 5
	fmt.Print("- ")
	for i := range 5 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func loopOverEvenNumbers() {
	fmt.Println("\nEven For Loop Example:")
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
	fmt.Println("\nWhile Loop Example:")
	fmt.Print("- ")
	i := 1
	for i <= 10 {
		fmt.Printf("%d ", i)
		i *= 2
	}
	fmt.Println()
}

func labelsExample() {
	fmt.Println("\nLabels: break/continue outer")
outer:
	for i := 0; i <= 3; i++ { // keep the grid tiny to see the effect
		for j := 0; j <= 3; j++ {
			fmt.Printf("  visiting (i=%d, j=%d)\n", i, j)

			// Skip to next i when sum is 2
			if i+j == 2 {
				fmt.Println("  -> continue outer (i+j==2)")
				continue outer
			}

			// Break everything once we reach (3,0)
			if i == 3 && j == 0 {
				fmt.Println("  -> break outer")
				break outer
			}
		}
	}
}

func gotoExample() {
	fmt.Println("\ngoto (early exit demo)")
	// Simple example: scan small slice, jump to label when found.
	target := 9
	for _, n := range []int{2, 4, 6, 8, 9, 10} {
		if n == target {
			goto Found
		}
	}
	fmt.Println("not found")
	return

Found:
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

		if absolute(z-prev) < THRESHOLD {
			break // Exit loop
		}
	}
	return z, iterations
}

func absolute(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	backwardForExample()
	forwardForExample()
	loopOverEvenNumbers()
	whileLoopExample()
	labelsExample()
	gotoExample()

	fmt.Println("\nSquare root calculation using Newton's method:")
	result, iterations := sqrt(11)
	fmt.Printf("Square root of 11: %.2f (found in %d iterations)\n", result, iterations)
}
