package main

import (
	"fmt"
)

// Define a struct (like a class in other languages)
type Point struct {
	Lat float64 // if field name is capitalized, it's exported (public)
	Lng float64
}

var locations = map[Point]string{
	{Lat: 40.785091, Lng: -73.968285}: "Central Park",
	{Lat: 48.858844, Lng: 2.294351}:   "Eiffel Tower",
	{Lat: 43.7228, Lng: 10.4018}:      "Pisa Tower",
	{Lat: 29.9792, Lng: 31.1342}:      "Egyptian Pyramids",
}

type Person struct {
	Name     string
	Age      int
	Location Point // Nested struct
}

// Value receiver method - does not modify the original struct
func (p Person) Greet() string {
	loc, ok := locations[p.Location]
	if !ok {
		loc = fmt.Sprintf("Unknown location (%f, %f)", p.Location.Lat, p.Location.Lng)
	}

	return fmt.Sprintf("Hello, my name is %s, I am %d years old and I live in %s.", p.Name, p.Age, loc)
}

// Pointer receiver method - can modify the original struct
func (p *Person) UpdateAge(a int) {
	if p != nil && a >= 0 { // Check if pointer is not nil and age is valid
		p.Age = a
	}
}

// Constructor function for Person struct
// Used for small to medium sized structs (remains on the stack)
func newPerson(n string, a int, lat, lng float64) Person {
	return Person{n, a, Point{Lat: lat, Lng: lng}} // Positional fields
}

// Constructor function that returns a pointer to Person struct
// Useful for larger structs to avoid copying (but saves the struct on the heap)
func newPersonPointer(n string, a int, lat, lng float64) *Person {
	return &Person{Name: n, Age: a, Location: Point{Lat: lat, Lng: lng}} // Named fields
}

func main() {
	nyCentralPark := Point{Lat: 40.785091, Lng: -73.968285} // Named fields
	eiffelTower := Point{48.858844, 2.294351}               // Positional fields

	fmt.Printf("Central Park Point: (%f, %f) and Eiffel Tower Point: (%f, %f)\n", nyCentralPark.Lat, nyCentralPark.Lng, eiffelTower.Lat, eiffelTower.Lng)

	// Anonymous struct
	msg := struct {
		Message string
		From    Point
	}{"Hello America it's me Mario!", Point{43.7228, 10.4018}}

	fmt.Printf("Message: %s, From: (%f, %f)\n", msg.Message, msg.From.Lat, msg.From.Lng)

	p1 := newPerson("Egyptian", 100, 29.9792, 31.1342)
	fmt.Println(p1.Greet())

	// You don't need to use `(&p1).UpdateAge()`
	// since Go automatically dereferences the pointer when calling methods
	p1.UpdateAge(101)
	fmt.Println("After update:", p1.Greet())

	p2 := newPersonPointer("French", 50, 48.858844, 2.294351)
	fmt.Println(p2.Greet())
	p2.UpdateAge(51)
	fmt.Println("After update:", p2.Greet())

	p3 := new(Person) // This creates a new Person with zero values
	p4 := &Person{}   // This also creates a new Person with zero values
	fmt.Println(p3.Greet())
	fmt.Println(p4.Greet())
}
