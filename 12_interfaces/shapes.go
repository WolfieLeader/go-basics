package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
	Circumference() float64
}

// Define a Circle type
type Circle struct{ Radius float64 }

// Implement Shape interface for Circle
func (c Circle) Area() float64          { return 3.14 * c.Radius * c.Radius }
func (c Circle) Circumference() float64 { return 2 * 3.14 * c.Radius }

// Define a Rectangle type
type Rectangle struct{ Width, Height float64 }

// Implement Shape interface for Rectangle Pointer
func (r *Rectangle) Area() float64          { return r.Width * r.Height }
func (r *Rectangle) Circumference() float64 { return 2 * (r.Width + r.Height) }

// Define a Line type as a type alias for float6s
type Line float64

// If we would make a type alias like so:
// type Line = float64
// Then we could not implement methods on it, as type aliases do not allow method sets.

// Implementing a method on Line type alias
func (l Line) Area() float64          { return 0.0 }
func (l Line) Circumference() float64 { return float64(l) }

// Function that takes a Shape interface
func printShape(s Shape) {
	fmt.Printf("Type: %T - Area: %.2f, Circumference: %.2f\n", s, s.Area(), s.Circumference())
}

func shapeExample() {
	fmt.Println("\nShape Example:")
	var sPtr Shape
	fmt.Printf("Type of sPtr: %T\n", sPtr) // p is nil, so it doesn't hold any value

	c, r, l := Circle{5.0}, Rectangle{3.0, 4.0}, Line(10.0)

	printShape(c)
	printShape(&r) // Pass rectangle as a pointer since the interface is implemented by a pointer receiver
	printShape(l)
}
