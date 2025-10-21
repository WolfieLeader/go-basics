package main

import (
	"encoding/base64"
	"fmt"
)

func base64Example() {
	data := []byte(HELLO_WORLD)
	fmt.Printf("- Original: %s\n", data)

	b64 := base64.StdEncoding.EncodeToString(data)    // Base64 uses + and /
	b64Url := base64.URLEncoding.EncodeToString(data) // Base64 URL uses - and _

	b64NoPad1 := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
	b64NoPad2 := base64.RawStdEncoding.EncodeToString(data)
	b64UrlNoPad1 := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
	b64UrlNoPad2 := base64.RawURLEncoding.EncodeToString(data)

	fmt.Printf("- Base64 Encoded(with padding): %s\n", b64)
	fmt.Printf("- Base64 URL Encoded(with padding): %s\n", b64Url)
	fmt.Printf("- Base64 Encoded(no padding): %s (ok=%t)\n", b64NoPad1, b64NoPad1 == b64NoPad2)
	fmt.Printf("- Base64 URL Encoded(no padding): %s (ok=%t)\n", b64UrlNoPad1, b64UrlNoPad1 == b64UrlNoPad2)
	fmt.Printf("- Base64 Encoded vs Base64 URL Encoded Are They Equal? %v\n", b64 == b64Url)
	fmt.Printf("- Base64 No Padding vs Base64 URL No Padding Are They Equal? %v\n", b64NoPad1 == b64UrlNoPad1)

	b64Decoded, _ := base64.StdEncoding.DecodeString(b64)
	b64UrlDecoded, _ := base64.URLEncoding.DecodeString(b64Url)

	b64NoPadDecoded1, _ := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(b64NoPad2)
	b64NoPadDecoded2, _ := base64.RawStdEncoding.DecodeString(b64NoPad2)
	b64UrlNoPadDecoded1, _ := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(b64UrlNoPad2)
	b64UrlNoPadDecoded2, _ := base64.RawURLEncoding.DecodeString(b64UrlNoPad2)

	fmt.Printf("- Are all decoded match original? %v\n", areAllEqual(b64Decoded, b64UrlDecoded, b64NoPadDecoded1, b64NoPadDecoded2, b64UrlNoPadDecoded1, b64UrlNoPadDecoded2, data))
}
