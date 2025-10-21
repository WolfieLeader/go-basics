package main

import (
	"encoding/hex"
	"fmt"
)

func hexExample() {
	data := []byte(HELLO_WORLD)
	fmt.Printf("- Original: %s\n", data)

	encoded := hex.EncodeToString(data)
	fmt.Printf("- Hex Encoded: %s\n", encoded)

	decoded, _ := hex.DecodeString(encoded)
	fmt.Printf("- Decoded matches original? %v\n", string(decoded) == string(data))
}
