package main

import "fmt"

func main() {
	// Maps are unordered collections of key-value pairs
	// They are similar to dictionaries in Python or hash tables in other languages
	studentGrades := map[string]int{
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
	studentGrades["Alice"] = 95 // Update Alice's grade

	// Accessing values using keys
	fmt.Printf("Alice's grade: %d\n", studentGrades["Alice"])
	name, ok := studentGrades["NonExistent"]
	fmt.Printf("NonExistent's grade: %d, exists: %t\n", name, ok)

	// Iterating over the map
	for key, value := range studentGrades {
		fmt.Printf("- %s has a grade of %d\n", key, value)
	}

	// Deleting a key-value pair
	delete(studentGrades, "Bob")
	println("After deleting Bob, the map contains:")
	for student, grade := range studentGrades {
		fmt.Printf("- %s has a grade of %d\n", student, grade)
	}
}
