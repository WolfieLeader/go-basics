package main

import "fmt"

type HttpError struct {
	StatusCode int
	Message    string
}

// Error interface is a built-in interface in Go
// It has a single method Error() string
func (e *HttpError) Error() string {
	if e == nil {
		return "(nil HttpError)"
	}

	return fmt.Sprintf("(HTTP %d: %s)", e.StatusCode, e.Message)
}

type IP [4]byte

// Stringer interface is a built-in interface in Go
// It has a single method String() string
func (ip IP) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func errorAndStringerExample() {
	err := &HttpError{404, "Not Found"}
	var nilErr *HttpError

	fmt.Printf("- Error: %s, nilErr: %s\n", err, nilErr) // This will call the Error() method

	hosts := map[string]IP{"localhost": {127, 0, 0, 1}, "google": {8, 8, 8, 8}, "cloudflare": {1, 1, 1, 1}}

	for k, v := range hosts {
		fmt.Printf("- %s IP: %s\n", k, v) // This will call the String() method
	}
}
