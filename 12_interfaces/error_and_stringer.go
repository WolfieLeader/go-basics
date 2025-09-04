package main

import "fmt"

// Error interface is a built-in interface in Go
// It has a single method Error() string

type HttpError struct {
	StatusCode int
	Message    string
}

// The Error method implements the error interface
func (e *HttpError) Error() string {
	if e == nil {
		return "(nil HttpError)"
	}

	return fmt.Sprintf("(HTTP %d: %s)", e.StatusCode, e.Message)
}

// Returns new custom error pointer since the method is defined on a pointer receiver
func getNotFound() *HttpError {
	// The garbage collector will put the HttpError in the heap
	return &HttpError{404, "Not Found"}
}

// Stringer interface is a built-in interface in Go
// It has a single method String() string

type IP [4]byte

func (ip IP) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func errorAndStringerExample() {
	fmt.Println("\nError and Stringer Example:")

	// Create a new HttpError
	err := getNotFound()
	var nilErr *HttpError
	fmt.Printf("Error: %s, nilErr: %s\n", err, nilErr) // This will call the Error() method

	hosts := map[string]IP{
		"localhost":  {127, 0, 0, 1},
		"google":     {8, 8, 8, 8},
		"cloudflare": {1, 1, 1, 1},
	}

	for k, v := range hosts {
		fmt.Printf("%s IP: %s\n", k, v) // This will call the String() method
	}
}
