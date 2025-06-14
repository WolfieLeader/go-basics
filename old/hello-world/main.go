package main

import (
	"fmt"
	// go get rsc.io/quote
	// or
	// import it and then run
	// go mod tidy
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
