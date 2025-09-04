package main

import "fmt"

func arrayExample1() {
	fmt.Println("\nArray Example:")

	// Fixed string array with length 3
	strArr := [3]string{"Hello", "World", "!"}
	fmt.Println("- Fixed array:", strArr)

	// Fixed integer array with length 5
	var intArr [5]int // [0 0 0 0 0]
	intArr[0] = 1     // [1 0 0 0 0]
	intArr[2] = 3     // [1 0 3 0 0]
	intArr[4] = 5     // [1 0 3 0 5]
	fmt.Println("- Fixed integer array:", intArr)
}

func arrayExample2() {
	fmt.Println("\nArray Example 2:")
	// `[...]T` is a shorthand for array with length inferred from the number of elements
	intArr := [...]int{100, 200, 300, 400, 500}

	// Going through the array using a for loop
	var prt string
	for index, value := range intArr {
		prt += fmt.Sprintf("intArr[%d] = %d, ", index, value)
	}
	fmt.Println("- ", prt[:len(prt)-2]) // Remove trailing comma and space
}
