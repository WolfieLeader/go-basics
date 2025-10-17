package main

import (
	"fmt"
	"maps"
)

func mapExample() {
	// Maps are unordered collections of key-value pairs
	// They are similar to dictionaries in Python or hash tables in other languages
	mathGrades := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 92,
	}

	mathGrades["David"] = 88 // Add a new key-value pair
	mathGrades["Alice"] = 95 // Update Alice's grade

	// Alternatively, you can create an empty map and add key-value pairs later
	scienceGrades := make(map[string]int)
	scienceGrades["Bob"], scienceGrades["Charlie"] = 72, 89

	// Accessing values using keys (returns zero value if key does not exist)
	fmt.Printf("Alice's math grade: %d, science grade (not available): %d\n", mathGrades["Alice"], scienceGrades["Alice"])

	name, ok := mathGrades["NonExistent"]
	fmt.Printf("NonExistent's math grade: %d, exists: %t\n", name, ok)

	// Iterating over the map
	for key, value := range mathGrades {
		fmt.Printf("- %s has a math grade of %d\n", key, value)
	}

	delete(mathGrades, "Bob") // Deleting a key-value pair

	fmt.Println("After deleting Bob, the map contains:")
	for student, grade := range mathGrades {
		fmt.Printf("- %s has a math grade of %d\n", student, grade)
	}
}

func mapsPackageExample() {
	langVersions := map[string]string{
		"Zig":    "0.15.2",
		"Go":     "1.25.3",
		"Python": "3.14.0",
	}

	cloned := maps.Clone(langVersions) // Returns a cloned copy
	fmt.Printf("- Cloned: %v, Src: %v, Are Equal? %t\n", cloned, langVersions, maps.Equal(langVersions, cloned))

	delete(cloned, "Python")
	fmt.Printf("- After deleting Python from cloned map: %v, Are Equal? %t\n", cloned, maps.Equal(langVersions, cloned))

	lowLevelLangs := map[string]string{
		"Go":   "1.18.0", // Old version
		"Rust": "1.90.0",
	}

	// Copies key-value pairs from source to destination
	// If a key exists in both maps, the value from the source overwrites the value in the destination
	maps.Copy(lowLevelLangs, cloned)
	fmt.Printf("- Low-level languages after copy: %v\n", lowLevelLangs)

	// `maps.Keys()` and `maps.Values()` return iterators which can be used in `for range` loops
	fmt.Print("- Keys from lowLevelLangs:")
	for key := range maps.Keys(lowLevelLangs) {
		fmt.Printf(" %s,", key)
	}
	fmt.Println()

	fmt.Print("- Values from lowLevelLangs:")
	for value := range maps.Values(lowLevelLangs) {
		fmt.Printf(" %s,", value)
	}
	fmt.Println()
}

func main() {
	fmt.Println("Map Example:")
	mapExample()
	fmt.Println()

	fmt.Println("Maps Package Example:")
	mapsPackageExample()
	fmt.Println()
}
