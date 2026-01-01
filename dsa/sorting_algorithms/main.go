package main

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/dsa/sorting_algorithms/sorting"
)

func main() {
	arr := []int{0, 64, 32, 50, 50, -5, 11, 90, 25, 15000, 200, 4, 3, 50, 2, -30, 23, 0, -1, 1, 50}

	fmt.Println(sorting.Bubble(arr), "- Bubble Sort")
	fmt.Println(sorting.Selection(arr), "- Selection Sort")
	fmt.Println(sorting.Insertion(arr), "- Insertion Sort")
	fmt.Println(sorting.Merge(arr), "- Merge Sort")
	fmt.Println(sorting.Quick(arr), "- Quick Sort")
	fmt.Println(sorting.QuickHoare(arr), "- Quick Hoare Sort")
	fmt.Println(sorting.Quick3Way(arr), "- Quick 3-Way Sort")
	fmt.Println(sorting.Bucket(arr), "- Bucket Sort")
	fmt.Println(sorting.Counting(arr), "- Counting Sort")
	fmt.Println(sorting.Radix(arr), "- Radix Sort")
	fmt.Println(sorting.Cocktail(arr), "- Cocktail Sort")
	fmt.Println(sorting.Shell(arr), "- Shell Sort")
	fmt.Println(sorting.Gnome(arr), "- Gnome Sort")
	fmt.Println(sorting.Comb(arr), "- Comb Sort")
	fmt.Println(sorting.Heap(arr), "- Heap Sort")
}
