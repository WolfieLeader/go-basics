package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

func toHex(data []byte, buf *bytes.Buffer) string {
	encoder := hex.NewEncoder(buf)
	encoder.Write(data)
	return buf.String()
}

func fromHex(buf *bytes.Buffer) string {
	decoder := hex.NewDecoder(buf)
	decoded, _ := io.ReadAll(decoder)
	return string(decoded)
}

func hexEncoderExample() {
	data := []byte(JAPANESE)
	fmt.Printf("- Original: %s\n", data)

	var hexBuf bytes.Buffer
	encoded := toHex(data, &hexBuf)
	fmt.Printf("- Hex Encoded: %s\n", encoded)

	decoded := fromHex(&hexBuf)
	fmt.Printf("- Decoded matches original? %v\n", decoded == string(data))
}
