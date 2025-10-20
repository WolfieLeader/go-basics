package encoding

import (
	"bytes"
	"encoding/base32"
	"fmt"
	"io"
)

func Base32Example() {
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

func toBase32(data []byte, encoding *base32.Encoding, buf *bytes.Buffer) string {
	encoder := base32.NewEncoder(encoding, buf)
	encoder.Write(data)
	encoder.Close()
	return buf.String()
}

func fromBase32(encoding *base32.Encoding, buf *bytes.Buffer) []byte {
	decoder := base32.NewDecoder(encoding, buf)
	decoded, _ := io.ReadAll(decoder)
	return decoded
}

func Base32EncoderExample() {
	data := []byte(JAPANESE)
	fmt.Printf("- Original: %s\n", data)

	var b32Buf, b32HexBuf, b32NoPadBuf, b32HexNoPadBuf bytes.Buffer
	b32 := toBase32(data, base32.StdEncoding, &b32Buf)
	b32Hex := toBase32(data, base32.HexEncoding, &b32HexBuf)

	b32NoPad := toBase32(data, base32.StdEncoding.WithPadding(base32.NoPadding), &b32NoPadBuf)
	b32HexNoPad := toBase32(data, base32.HexEncoding.WithPadding(base32.NoPadding), &b32HexNoPadBuf)

	fmt.Printf("- Base32 Encoded(with padding): %s\n", b32)
	fmt.Printf("- Base32 Hex Encoded(with padding): %s\n", b32Hex)
	fmt.Printf("- Base32 Encoded(no padding): %s\n", b32NoPad)
	fmt.Printf("- Base32 Hex Encoded(no padding): %s\n", b32HexNoPad)
	fmt.Printf("- Base32 Encoded vs Base32 Hex Encoded Are They Equal? %t\n", b32 == b32Hex)
	fmt.Printf("- Base32 No Padding vs Base32 Hex No Padding Are They Equal? %t\n", b32NoPad == b32HexNoPad)

	b32Decoded := fromBase32(base32.StdEncoding, &b32Buf)
	b32HexDecoded := fromBase32(base32.HexEncoding, &b32HexBuf)

	b32NoPadDecoded := fromBase32(base32.StdEncoding.WithPadding(base32.NoPadding), &b32NoPadBuf)
	b32HexNoPadDecoded := fromBase32(base32.HexEncoding.WithPadding(base32.NoPadding), &b32HexNoPadBuf)
	fmt.Printf("- Are all decoded match original? %v\n", areAllEqual(b32Decoded, b32HexDecoded, b32NoPadDecoded, b32HexNoPadDecoded, data))
}
