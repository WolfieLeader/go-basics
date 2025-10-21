package main

import (
	"encoding/base32"
	"fmt"
)

func base32Example() {
	data := []byte(HELLO_WORLD)
	fmt.Printf("- Original: %s\n", data)

	b32 := base32.StdEncoding.EncodeToString(data)
	b32Hex := base32.HexEncoding.EncodeToString(data)

	b32NoPad := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(data)
	b32HexNoPad := base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(data)

	fmt.Printf("- Base32 Encoded(with padding): %s\n", b32)
	fmt.Printf("- Base32 Hex Encoded(with padding): %s\n", b32Hex)
	fmt.Printf("- Base32 Encoded(no padding): %s\n", b32NoPad)
	fmt.Printf("- Base32 Hex Encoded(no padding): %s\n", b32HexNoPad)
	fmt.Printf("- Base32 Encoded vs Base32 Hex Encoded Are They Equal? %t\n", b32 == b32Hex)
	fmt.Printf("- Base32 No Padding vs Base32 Hex No Padding Are They Equal? %t\n", b32NoPad == b32HexNoPad)

	b32Decoded, _ := base32.StdEncoding.DecodeString(b32)
	b32HexDecoded, _ := base32.HexEncoding.DecodeString(b32Hex)

	b32NoPadDecoded, _ := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(b32NoPad)
	b32HexNoPadDecoded, _ := base32.HexEncoding.WithPadding(base32.NoPadding).DecodeString(b32HexNoPad)

	fmt.Printf("- Are all decoded match original? %v\n", areAllEqual(b32Decoded, b32HexDecoded, b32NoPadDecoded, b32HexNoPadDecoded, data))
}
