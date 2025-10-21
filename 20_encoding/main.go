package main

import (
	"bytes"
	"fmt"
)

const (
	HELLO_WORLD = "ĤèlĬϴ ₩órł⫒🌍"
	JAPANESE    = "日本語🏯"
)

func areAllEqual(data ...[]byte) bool {
	if len(data) < 2 {
		return true
	}

	for i := 1; i < len(data); i++ {
		if !bytes.Equal(data[0], data[i]) {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println("Base64 Simple Example:")
	base64Example()
	fmt.Println()

	fmt.Println("Base64 Encoder Example:")
	base64EncoderExample()
	fmt.Println()

	fmt.Println("Hex Simple Example:")
	hexExample()
	fmt.Println()

	fmt.Println("Hex Encoder Example:")
	hexEncoderExample()
	fmt.Println()

	fmt.Println("Base32 Simple Example:")
	base32Example()
	fmt.Println()

	fmt.Println("Base32 Encoder Example:")
	base32EncoderExample()
	fmt.Println()
}
