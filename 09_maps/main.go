package main

import (
	"fmt"
	"maps"
)

func main() {
	// Maps are unordered collections of key-value pairs
	// They are similar to dictionaries in Python or hash tables in other languages
	mathGrades := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 92,
	}

	// Alternatively, you can create an empty map and add key-value pairs later
	versions := make(map[string]float64)

	// Adding key-value pairs to the map
	versions["Go"] = 1.24
	versions["Python"] = 3.14

	//Changing a value in the map
	mathGrades["Alice"] = 95

	// Accessing values using keys
	fmt.Printf("Alice's grade: %d\n", mathGrades["Alice"])

	name, ok := mathGrades["NonExistent"]
	fmt.Printf("NonExistent's grade: %d, exists: %t\n", name, ok)

	// Iterating over the map
	for key, value := range mathGrades {
		fmt.Printf("- %s has a grade of %d\n", key, value)
	}

	// Deleting a key-value pair
	delete(mathGrades, "Bob")
	fmt.Println("After deleting Bob, the map contains:")
	for student, grade := range mathGrades {
		fmt.Printf("- %s has a grade of %d\n", student, grade)
	}

	// Returns a cloned copy, instead of using `copy()` and pre-allocating
	cloned := maps.Clone(mathGrades)
	fmt.Printf("- Cloned: %v, Src: %v, Are Equal? %t\n", cloned, mathGrades, maps.Equal(mathGrades, cloned))

	sportGrades := map[string]int{
		"Alice": 25,
		"David": 30,
	}

	newSportGrades := maps.Clone(sportGrades)

	// Copies key-value pairs from source map to destination map
	// If a key exists in both maps, the value from the source map overwrites the value in the destination map
	maps.Copy(newSportGrades, mathGrades)
	fmt.Printf("- Old Sport Grades: %v, New: %v, Are Equal? %t\n", sportGrades, newSportGrades, maps.Equal(sportGrades, newSportGrades))

	// `maps.Keys()` and `maps.Values()` return iterators
	// Iterators will be covered later in detail

	fmt.Print("- Iterating over keys using maps.Keys():")
	for key := range maps.Keys(mathGrades) {
		fmt.Printf(" %s,", key)
	}
	fmt.Println()

	fmt.Print("- Iterating over values using maps.Values():")
	for value := range maps.Values(mathGrades) {
		fmt.Printf(" %d,", value)
	}
	fmt.Println()
}
