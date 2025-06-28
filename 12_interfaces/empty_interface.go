package main

import "fmt"

func typeAssertion(x interface{}) string {
	// the `type` keyword is used to switch on the type of the interface
	switch v := x.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float"
	case complex128:
		return "complex"
	default:
		return fmt.Sprintf("unknown type %T", v)
	}
}

func emptyInterfaceExample() {
	fmt.Println("\nEmpty Interface Example:")

	// Empty interface can hold any type
	var anything interface{}

	anything = struct{ X string }{"Some String"}
	fmt.Printf("Type of anything: %T, Value: %v\n", anything, anything)

	anything = 10

	// Type assertion to convert interface{} back to int
	i := anything.(int) // This will panic if anything is not an int
	fmt.Printf("Type of i: %T, Value: %d\n", i, i)

	f, ok := anything.(float64) // This will not panic, ok will be false
	fmt.Printf("Type of f: %T, Value: %.2f(zero value), Ok: %t\n", f, f, ok)

	anything = 3 + 4i // complex number
	t := typeAssertion(anything)
	fmt.Printf("Type of anything: %T, Value: %v, Type Assertion Result: %s\n", anything, anything, t)
}
