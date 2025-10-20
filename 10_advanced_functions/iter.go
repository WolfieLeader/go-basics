package main

import (
	"fmt"
	"iter"
	"strings"
)

// Returns an iterator that applies the given function to each element of the input array.
func transformNumbers(numbers []int, fn func(num int) int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, value := range numbers {
			// Apply the transformation function and yield the result
			if !yield(fn(value)) {
				return
			}
		}
	}
}

func iterSeq1Example() {
	numbers := []int{1, 2, 3, 4, 5}
	squareFn := func(num int) int { return num * num }

	squared := transformNumbers(numbers, squareFn)
	fmt.Println("- Original numbers:", numbers)

	//There are two ways to consume the iterator:

	// 1. Using the `range` keyword which is a sugar syntax for iterators
	fmt.Print("- Squared numbers using range: ")
	for num := range squared {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 2. Using the function form (this is what `range` uses under the hood)
	fmt.Print("- Squared numbers using function form: ")
	squared(func(num int) bool {
		fmt.Printf("%d ", num)
		return true // Continue iteration
	})
	fmt.Println()
}

func transformList(list map[string]int, fn func(key string, value int) (string, int)) iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for key, value := range list {
			newKey, newValue := fn(key, value)
			if !yield(newKey, newValue) {
				return
			}
		}
	}
}

func iterSeq2Example() {
	doctorRatings := map[string]int{
		"John Smith":  5,
		"Jane Doe":    4,
		"Emily Davis": 3,
	}
	rateScoreFn := func(fullName string, rating int) (string, int) {
		fields := strings.Fields(fullName)
		title := "Dr. " + fields[len(fields)-1]
		outOf100 := rating * 20
		return title, outOf100
	}

	doctorScores := transformList(doctorRatings, rateScoreFn)
	fmt.Printf("- Before Transformation: %v\n", doctorRatings)

	fmt.Print("- Doctor Scores using range:")
	for name, score := range doctorScores {
		fmt.Printf("  %s: %d,", name, score)
	}
	fmt.Println()

	fmt.Print("- Doctor Scores using function form:")
	doctorScores(func(name string, score int) bool { fmt.Printf("  %s: %d,", name, score); return true })
	fmt.Println()
}
