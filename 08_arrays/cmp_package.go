package main

import (
	"cmp"
	"fmt"
)

// "cmp" package provides utility functions to compare values
// it's mostly used for generics
func cmpPackageExample() {
	var a string // Zero value of string
	b := ""      // Another zero value of string
	c := "ABC"

	// `cmp.Or` useful for fallbacks
	fmt.Printf("- First non-zero value among a, b and c is: %q\n", cmp.Or(a, b, c))

	b = "DEF"
	fmt.Printf("- After setting b, first non-zero value among a, b and c is: %q\n", cmp.Or(a, b, c))

	fmt.Printf("- Comparing 3 and 10: %d\n", cmp.Compare(3, 10))   // -1 because 3 < 10
	fmt.Printf("- Comparing 1 and 1: %d\n", cmp.Compare(1, 1))     // 0 because 1 == 1
	fmt.Printf("- Comparing 15 and 20: %d\n", cmp.Compare(15, 20)) // -1 because 15 < 20

	fmt.Printf("- Is ABC < DEF? %t\n", cmp.Less("ABC", "DEF"))
}
