package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// ? This is good practice to use pointer receivers
// ? it will automatically do (&v) for you
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Shape interface {
	Area() float64
	Circumference() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r *Rectangle) Circumference() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Type: %T - Area: %0.2f, Circumference: %0.2f\n", s, s.Area(), s.Circumference())
}

func EmptyInterface() {
	var i interface{} //? Empty interface can hold any type

	i = 23
	fmt.Printf("i is of type %T with value %v\n", i, i)

	i = "Hello World"
	fmt.Printf("i is now of type %T with value %v\n", i, i)

	i = Vertex{9, 1}
	fmt.Printf("i is now of type %T with value %v\n", i, i)

	i = MyFloat(3.14)
	fmt.Printf("i is now of type %T with value %v\n", i, i)
}

func EmptyInterfaceTypeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) //? panic because there is no ok
	fmt.Println(f)
}

func TypeSwitch(i interface{}) {
	switch v := i.(type) { //? type keyword is used to switch on the type of the interface
	case int:
		fmt.Printf("Type is int with value %d\n", v)
	case string:
		fmt.Printf("Type is string with value %s\n", v)
	case float64:
		fmt.Printf("Type is float64 with value %f\n", v)
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}

func (v Vertex) String() string {
	//? This is a Stringer method, which allows us to define how the Vertex type should be represented as a string.
	return fmt.Sprintf("Vertex(X: %0.2f, Y: %0.2f)", v.X, v.Y)
}
