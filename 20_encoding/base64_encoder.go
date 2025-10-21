package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

func toBase64(data []byte, encoding *base64.Encoding, buf *bytes.Buffer) string {
	encoder := base64.NewEncoder(encoding, buf)
	encoder.Write(data)
	encoder.Close()
	return buf.String()
}

func fromBase64(encoding *base64.Encoding, buf *bytes.Buffer) []byte {
	decoder := base64.NewDecoder(encoding, buf)
	decoded, _ := io.ReadAll(decoder)
	return decoded
}

func base64EncoderExample() {
	data := []byte(JAPANESE)
	fmt.Printf("- Original: %s\n", data)

	var b64Buf, b64UrlBuf, b64NoPadBuf1, b64NoPadBuf2, b64UrlNoPadBuf1, b64UrlNoPadBuf2 bytes.Buffer
	b64 := toBase64(data, base64.StdEncoding, &b64Buf)
	b64Url := toBase64(data, base64.URLEncoding, &b64UrlBuf)

	b64NoPad1 := toBase64(data, base64.StdEncoding.WithPadding(base64.NoPadding), &b64NoPadBuf1)
	b64NoPad2 := toBase64(data, base64.RawStdEncoding, &b64NoPadBuf2)
	b64UrlNoPad1 := toBase64(data, base64.URLEncoding.WithPadding(base64.NoPadding), &b64UrlNoPadBuf1)
	b64UrlNoPad2 := toBase64(data, base64.RawURLEncoding, &b64UrlNoPadBuf2)

	fmt.Printf("- Base64 Encoded(with padding): %s\n", b64)
	fmt.Printf("- Base64 URL Encoded(with padding): %s\n", b64Url)
	fmt.Printf("- Base64 Encoded(no padding): %s (ok=%t)\n", b64NoPad1, b64NoPad1 == b64NoPad2)
	fmt.Printf("- Base64 URL Encoded(no padding): %s (ok=%t)\n", b64UrlNoPad1, b64UrlNoPad1 == b64UrlNoPad2)
	fmt.Printf("- Base64 Encoded vs Base64 URL Encoded Are They Equal? %v\n", b64 == b64Url)
	fmt.Printf("- Base64 No Padding vs Base64 URL No Padding Are They Equal? %v\n", b64NoPad1 == b64UrlNoPad1)

	b64Decoded := fromBase64(base64.StdEncoding, &b64Buf)
	b64UrlDecoded := fromBase64(base64.URLEncoding, &b64UrlBuf)

	b64NoPadDecoded1 := fromBase64(base64.StdEncoding.WithPadding(base64.NoPadding), &b64NoPadBuf1)
	b64NoPadDecoded2 := fromBase64(base64.RawStdEncoding, &b64NoPadBuf2)
	b64UrlNoPadDecoded1 := fromBase64(base64.URLEncoding.WithPadding(base64.NoPadding), &b64UrlNoPadBuf1)
	b64UrlNoPadDecoded2 := fromBase64(base64.RawURLEncoding, &b64UrlNoPadBuf2)
	fmt.Printf("- Are all decoded match original? %v\n", areAllEqual(b64Decoded, b64UrlDecoded, b64NoPadDecoded1, b64NoPadDecoded2, b64UrlNoPadDecoded1, b64UrlNoPadDecoded2, data))
}
