package main

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/19_data_manipulation/encoding"
)

func main() {
	fmt.Println("Base64 Simple Example:")
	encoding.Base64Example()
	fmt.Println()

	fmt.Println("Base64 Encoder Example:")
	encoding.Base64EncoderExample()
	fmt.Println()

	fmt.Println("Hex Simple Example:")
	encoding.HexExample()
	fmt.Println()

	fmt.Println("Hex Encoder Example:")
	encoding.HexEncoderExample()
	fmt.Println()

	fmt.Println("Base32 Simple Example:")
	encoding.Base32Example()
	fmt.Println()

	fmt.Println("Base32 Encoder Example:")
	encoding.Base32EncoderExample()
	fmt.Println()
}
