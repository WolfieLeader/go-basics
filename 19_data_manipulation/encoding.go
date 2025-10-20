package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	helloWorld = "ĤèlĬϴ ₩órł⫒🌍"
	japanese   = "日本語🏯"
)

func base64SimpleExample() {
	data := []byte(helloWorld)
	fmt.Printf("- Original: %s\n", data)

	b64 := base64.StdEncoding.EncodeToString(data) // Base64 uses + and /
	b64NoPad := base64.RawStdEncoding.EncodeToString(data)
	b64Url := base64.URLEncoding.EncodeToString(data) // Base64 URL uses - and _
	b64UrlNoPad := base64.RawURLEncoding.EncodeToString(data)
	fmt.Printf("- Base64 Encoded(with padding): %s\n- Base64 Encoded(no padding): %s\n- Base64 URL Encoded(with padding): %s\n- Base64 URL Encoded(no padding): %s\n- Are they equal? %v\n", b64, b64NoPad, b64Url, b64UrlNoPad, b64 == b64Url && b64NoPad == b64UrlNoPad && b64 != b64NoPad)

	b64Decoded, _ := base64.StdEncoding.DecodeString(b64)
	b64UrlDecoded, _ := base64.URLEncoding.DecodeString(b64Url)
	b64NoPadDecoded, _ := base64.RawStdEncoding.DecodeString(b64NoPad)
	b64UrlNoPadDecoded, _ := base64.RawURLEncoding.DecodeString(b64UrlNoPad)
	fmt.Printf("- Are all decoded match original? %v\n", string(b64Decoded) == string(b64UrlDecoded) && string(b64NoPadDecoded) == string(b64UrlNoPadDecoded) && string(b64Decoded) == string(b64NoPadDecoded) && string(b64Decoded) == string(data))
}

func toBase64(data []byte, encoding *base64.Encoding, buf *bytes.Buffer) string {
	encoder := base64.NewEncoder(encoding, buf)
	encoder.Write(data)
	encoder.Close()
	return buf.String()
}

func fromBase64(encoding *base64.Encoding, buf *bytes.Buffer) string {
	decoder := base64.NewDecoder(encoding, buf)
	decoded, _ := io.ReadAll(decoder)
	return string(decoded)
}

func base64EncoderExample() {
	data := []byte(japanese)
	fmt.Printf("- Original: %s\n", data)

	var b64Buf, b64UrlBuf, b64NoPadBuf, b64UrlNoPadBuf bytes.Buffer
	b64 := toBase64(data, base64.StdEncoding, &b64Buf)
	b64Url := toBase64(data, base64.URLEncoding, &b64UrlBuf)
	b64NoPad := toBase64(data, base64.RawStdEncoding, &b64NoPadBuf)
	b64UrlNoPad := toBase64(data, base64.RawURLEncoding, &b64UrlNoPadBuf)
	fmt.Printf("- Base64 Encoded(with padding): %s\n- Base64 Encoded(no padding): %s\n- Base64 URL Encoded(with padding): %s\n- Base64 URL Encoded(no padding): %s\n- Are they equal? %v\n", b64, b64NoPad, b64Url, b64UrlNoPad, b64 == b64Url && b64NoPad == b64UrlNoPad && b64 != b64NoPad)

	decodedB64 := fromBase64(base64.StdEncoding, &b64Buf)
	decodedB64Url := fromBase64(base64.URLEncoding, &b64UrlBuf)
	decodedB64NoPad := fromBase64(base64.RawStdEncoding, &b64NoPadBuf)
	decodedB64UrlNoPad := fromBase64(base64.RawURLEncoding, &b64UrlNoPadBuf)
	fmt.Printf("- Are all decoded match original? %v\n", decodedB64 == decodedB64Url && decodedB64NoPad == decodedB64UrlNoPad && decodedB64 == decodedB64NoPad && decodedB64 == string(data))
}

func hexSimpleExample() {
	data := []byte(helloWorld)
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

func hexEncoderExample() {
	data := []byte(japanese)
	fmt.Printf("- Original: %s\n", data)

	var hexBuf bytes.Buffer
	encoded := toHex(data, &hexBuf)
	fmt.Printf("- Hex Encoded: %s\n", encoded)

	decoded := fromHex(&hexBuf)
	fmt.Printf("- Decoded matches original? %v\n", decoded == string(data))
}
