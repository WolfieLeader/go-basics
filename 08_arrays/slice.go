package main

import "fmt"

func sliceExample1() {
	floatSlice := []float64{1.1, 2.2} // Slice is a dynamically-sized array

	// `len` gives the current number of elements,
	// `cap` gives the max elements the slice can hold without reallocating
	floatSlice[0] = 0.1 // [0.1, 2.2]
	fmt.Printf("- Slice: %v, Length: %d, Capacity: %d\n", floatSlice, len(floatSlice), cap(floatSlice))

	// When capacity is exceeded, the capacity is doubled in size
	floatSlice = append(floatSlice, 3.3) // [0.1, 2.2, 3.3]
	fmt.Printf("- Slice: %v, Length: %d, Capacity: %d\n", floatSlice, len(floatSlice), cap(floatSlice))

	floatSlice = append(floatSlice, 4.4, 5.5) // [0.1, 2.2, 3.3, 4.4, 5.5]
	fmt.Printf("- Slice: %v, Length: %d, Capacity: %d\n", floatSlice, len(floatSlice), cap(floatSlice))
}

func sliceExample2() {
	fixedArray := [4]string{"John", "Paul", "George", "Ringo"}

	//`make` keyword creates a dynamic type
	slice1 := make([]string, 2)    // len 2 and cap 2
	slice2 := make([]string, 0, 2) // len 0 and cap 2

	slice1 = fixedArray[0:2] // Take first two elements
	slice2 = fixedArray[1:4] // Take from index 1 to 3 (4 is excluded)

	// Slices are references to underlying arrays
	// Modifying a slice modifies the underlying array
	slice2[0] = "X"

	fmt.Printf("- Fixed array: %v, Length: %d\n", fixedArray, len(fixedArray))
	fmt.Printf("- Slice 1 (0 and 1): %v, Length: %d, Capacity: %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("- Slice 2 (1, 2 and 3): %v, Length: %d, Capacity: %d\n", slice2, len(slice2), cap(slice2))
}

func sliceCopyExample() {
	src := []byte{'A', 'B', 'C'}

	dst := make([]byte, len(src))        // Allocate new slice of same length
	smallDst := make([]byte, len(src)-1) // Allocate smaller slice

	n1 := copy(dst, src) // n is number of elements copied
	n2 := copy(smallDst, src)

	dst[0] = 'Y' // Modify target to show independence from source

	fmt.Printf("- Source slice: %v\n", src)
	fmt.Printf("- Destination slice: %v, copied %d elements\n", dst, n1)
	fmt.Printf("- Smaller destination slice: %v, copied %d elements\n", smallDst, n2)
}
