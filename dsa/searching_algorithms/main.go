package main

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/dsa/searching_algorithms/searching"
)

func main() {
	unsortedArray := []int{50, 20, 0, 10, 60, 40, 80, 70, 30}
	sortedArray := []int{0, 10, 20, 30, 40, 50, 60, 70, 80}

	target := 30
	index := searching.LinearSearch(unsortedArray, target)
	if index != -1 {
		fmt.Println("Found", target, "at index", index)
	} else {
		fmt.Println(target, "not found")
	}

	index = searching.BinarySearch(sortedArray, target)
	if index != -1 {
		fmt.Println("Found", target, "at index", index)
	} else {
		fmt.Println(target, "not found")
	}
}
