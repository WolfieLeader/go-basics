package main

import "fmt"

func arrayExample1() {
	fixedStrArray := [3]string{"ABC", "DEF", "GHI"}
	fmt.Println("- Fixed array of 3 string elements:", fixedStrArray)

	var fixedUintArray [5]uint8 // [0 0 0 0 0]
	fixedUintArray[0] = 1       // [1 0 0 0 0]
	fixedUintArray[2] = 3       // [1 0 3 0 0]
	fixedUintArray[4] = 5       // [1 0 3 0 5]
	fmt.Println("- Fixed array of 5 uint8 elements:", fixedUintArray)
}

func arrayExample2() {
	inferredIntArray := [...]int{10, 8, 6, 4, 2} // `...` infers the length

	totalWithIndex := 0
	for index, value := range inferredIntArray {
		totalWithIndex += value * (index + 1) // Just to use the index in some way
	}
	fmt.Printf("- Inferred array: %v, Total with index: %d\n", inferredIntArray, totalWithIndex)
}
