package encoding

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

func HexExample() {
	data := []byte(HELLO_WORLD)
	fmt.Printf("- Original: %s\n", data)

	encoded := hex.EncodeToString(data)
	fmt.Printf("- Hex Encoded: %s\n", encoded)

	decoded, _ := hex.DecodeString(encoded)
	fmt.Printf("- Decoded matches original? %v\n", string(decoded) == string(data))
}

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

func HexEncoderExample() {
	data := []byte(JAPANESE)
	fmt.Printf("- Original: %s\n", data)

	var hexBuf bytes.Buffer
	encoded := toHex(data, &hexBuf)
	fmt.Printf("- Hex Encoded: %s\n", encoded)

	decoded := fromHex(&hexBuf)
	fmt.Printf("- Decoded matches original? %v\n", decoded == string(data))
}
