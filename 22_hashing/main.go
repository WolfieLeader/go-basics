package main

import (
	"crypto/sha256"
	"fmt"
)

func sha256Example() {
	data := []byte("ĤèlĬϴ ₩órł⫒🌍")
	fmt.Printf("- Original: %s\n", data)

	hash := sha256.New()
	hash.Write(data)

	hashedData := hash.Sum(nil)
	fmt.Printf("- SHA-256 Hash: %x\n", hashedData)
}

func main() {
	fmt.Println("SHA-256 Hashing Example:")
	sha256Example()
	fmt.Println()
}
