package main

import (
	"errors"
	"fmt"
)

var locations = map[Point]string{
	{Lat: 40.785091, Lng: -73.968285}: "Central Park",
	{Lat: 48.858844, Lng: 2.294351}:   "Eiffel Tower",
	{Lat: 43.7228, Lng: 10.4018}:      "Pisa Tower",
	{Lat: 29.9792, Lng: 31.1342}:      "Egyptian Pyramids",
}

// Define a struct (like a class in other languages)
type Point struct {
	// The same rule: if capitalized, it is exported
	Lat float64
	Lng float64
}

type Person struct {
	Name     string
	Age      int
	Location Point // Nested struct
}

// This is a constructor-like function
func newPerson(n string, a int, lat, lng float64) Person {
	return Person{n, a, Point{Lat: lat, Lng: lng}}
}

// Method on Person type
func (p Person) Greet() string {
	loc, ok := locations[p.Location]
	if !ok {
		loc = fmt.Sprintf("Unknown location (%f, %f)", p.Location.Lat, p.Location.Lng)
	}

	return fmt.Sprintf("Hello, my name is %s, I am %d years old and I live in %s.", p.Name, p.Age, loc)
}

// Method with pointer receiver
// This allows us to modify the Person's age
func (p *Person) UpdateAge(a int) error {
	// important: always check for nil pointers before dereferencing them
	if p == nil {
		return errors.New("nil pointer dereference")
	}

	p.Age = a
	return nil
}

func newPersonPointer(n string, a int, lat, lng float64) *Person {
	// This returns a pointer to a Person
	return &Person{
		Name:     n,
		Age:      a,
		Location: Point{Lat: lat, Lng: lng},
	}
}

func main() {
	// Named fields
	nyCentralPark := Point{Lat: 40.785091, Lng: -73.968285}
	// Positional fields
	eiffelTower := Point{48.858844, 2.294351}

	fmt.Printf("Central Park Point: (%f, %f) and Eiffel Tower Point: (%f, %f)\n", nyCentralPark.Lat, nyCentralPark.Lng, eiffelTower.Lat, eiffelTower.Lng)

	// Anonymous struct
	msg := struct {
		Message string
		From    Point
	}{
		"Hello America it's me Mario!", Point{43.7228, 10.4018},
	}
	fmt.Printf("Message: %s, From: (%f, %f)\n", msg.Message, msg.From.Lat, msg.From.Lng)

	p1 := newPerson("Egyptian", 100, 29.9792, 31.1342)
	fmt.Println(p1.Greet())

	// You don't need to use (&p1).UpdateAge() since Go automatically dereferences the pointer when calling methods
	p1.UpdateAge(101)
	fmt.Println("After update:", p1.Greet())

	p2 := newPersonPointer("French", 50, 48.858844, 2.294351)
	fmt.Println(p2.Greet())
	p2.UpdateAge(51)
	fmt.Println("After update:", p2.Greet())

	p3 := new(Person) // This creates a new Person with zero values
	fmt.Println(p3.Greet())
}
