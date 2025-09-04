package main

import "fmt"

// Go Primitive Types:
// - string
// - bool (true, false)
// - int, int8, int16, int32, int64
// - uint, uint8, uint16, uint32, uint64 (positive integers)
// - float32, float64
// - complex64, complex128 (complex numbers)
// - byte (alias for uint8, represents a byte of data)
// - rune (alias for int32, represents a Unicode code point)

// Explicit type declaration
var firstVar string = "Hello World!"

// Type inference
var secondVar = 23

// Multiple variable declaration
var x, y = 0.75, 1.5
var (
	s string     = "Go!"
	b bool       //zero value is `false`
	i int        // zero value is `0`
	u uint       = 10
	f float64    = 1.2345
	c complex128 = 1 + 2i
)

// Constants
const pi = 3.14159265358979323846
const (
	maxInt64  int64  = (1 << 32)
	minInt64  int64  = -(1 << 32)
	maxUint64 uint64 = (1 << 64) - 1
	minUint64 uint64 = 0
)

const (
	// `iota` is a predeclared identifier that represents successive untyped integer constants
	Sunday = iota + 1 // 1 (iota starts at 0, so we add 1)
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	// `iota` can be used to create constants with a pattern
	_  = iota             // 0, ignored
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20 = 1048576
	GB                    // 1 << 30 = 1073741824
	TB                    // 1 << 40 = 1099511627776
)

func main() {
	// Short declaration (inside functions only)
	name := "John Doe"
	age := 25
	isEmployed := true

	// Print values
	fmt.Println("Name: "+name, " Age: ", age, " Employed: ", isEmployed)
	fmt.Printf(
		`Types:
    - String: %s
    - Boolean: %t
    - Int: %d, Converted to Float: %.1f
    - Uint: %d
    - Float: %f, Shortened: %.2f
    - Complex: %v
		- Type of Complex: %T`,
		s, b, i, float64(i), u, f, f, c, c)
}
