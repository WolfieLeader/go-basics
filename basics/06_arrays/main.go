package main

import "fmt"

func arrayExample() {
	fmt.Println("\nArray Example:")
	// Fixed string array with length 3
	strArr := [3]string{"Hello", "World", "!"}
	fmt.Println("Fixed array:", strArr)

	// Fixed integer array with length 5
	var intArr [5]int // [0 0 0 0 0]
	intArr[0] = 1     // [1 0 0 0 0]
	intArr[2] = 3     // [1 0 3 0 0]
	intArr[4] = 5     // [1 0 3 0 5]
	fmt.Println("Fixed integer array:", intArr)
}

func sliceExample1() {
	fmt.Println("\nSlice Example 1:")
	// Slice is a dynamically-sized array
	floatSlice := []float64{1.1, 2.2}
	floatSlice[0] = 0.1                       // Modify first element
	floatSlice = append(floatSlice, 3.3)      // Append to slice
	floatSlice = append(floatSlice, 4.4, 5.5) // Append multiple values
	fmt.Println("Slice:", floatSlice)

	// Length is the current number of elements in the slice
	// Capacity is the maximum number of elements the slice can hold without reallocating
	// When appending, if the capacity is exceeded, a new underlying array is created with double the capacity
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", floatSlice, len(floatSlice), cap(floatSlice))
}

func sliceExample2() {
	fmt.Println("\nSlice Example 2:")
	namesArr := [4]string{"John", "Paul", "George", "Ringo"}

	// Rule on thumb: make keyword creates a dynamic type
	strSlice1 := make([]string, 2)    // Create a slice with length 2 and capacity 2
	strSlice2 := make([]string, 0, 2) // Create a slice with length 0 and capacity 2

	// [init: end(exclusive)]
	strSlice1 = namesArr[0:2] // Slice from array, includes elements 0 and 1
	strSlice2 = namesArr[1:4] // Slice from array, includes elements 1, 2, and 3

	// Slices are references to underlying arrays
	strSlice2[0] = "X" // This will change the original array and the slices that reference it

	fmt.Printf("Names array: %v, Names length: %d\n", namesArr, len(namesArr))
	fmt.Printf("Slice 1: %v, Length: %d, Capacity: %d\n", strSlice1, len(strSlice1), cap(strSlice1))
	fmt.Printf("Slice 2: %v, Length: %d, Capacity: %d\n", strSlice2, len(strSlice2), cap(strSlice2))
}

func main() {
	arrayExample()
	sliceExample1()
	sliceExample2()
}
