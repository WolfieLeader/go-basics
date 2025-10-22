package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func copyExample() {
	src, err := os.Open("texts/lorem-ipsum.txt")
	if err != nil {
		log.Fatalf("Failed to open source file: %s", err)
	}
	defer src.Close()
	io.Copy(os.Stdout, src) // Copies from src to dst (os.Stdout here) until EOF
}

func copyNExample() {
	src, err := os.Open("texts/lorem-ipsum.txt")
	if err != nil {
		log.Fatalf("Failed to open source file: %s", err)
	}
	defer src.Close()
	io.CopyN(os.Stdout, src, 8) // Copies from src to dst (os.Stdout here) until EOF or n bytes read
}

func seekExample() {
	str := strings.NewReader("Seek sets the offset for the next Read or Write to offset")

	buf := make([]byte, 8)
	n, err := str.Read(buf)
	if err != nil {
		log.Fatalf("Read error: %s", err)
	}
	fmt.Printf("- First read (%d bytes): %q\n", n, buf[:n])

	newOffset, err := str.Seek(10, io.SeekStart)
	if err != nil {
		log.Fatalf("Seek error: %s", err)
	}
	fmt.Printf("- Moved offset to: %d (10 bytes from start)\n", newOffset)

	buf = make([]byte, 8)
	n, err = str.Read(buf)
	if err != nil {
		log.Fatalf("Read error: %s", err)
	}
	fmt.Printf("- After Seek, next read (%d bytes): %q\n", n, buf[:n])

	newOffset, err = str.Seek(-5, io.SeekCurrent)
	if err != nil {
		log.Fatalf("Seek error: %s", err)
	}
	fmt.Printf("- Moved offset back 5 bytes (relative to current): %d\n", newOffset)

	buf = make([]byte, 8)
	n, err = str.Read(buf)
	if err != nil {
		log.Fatalf("Read error: %s", err)
	}
	fmt.Printf("- Read after SeekCurrent (%d bytes): %q\n", n, buf[:n])

	newOffset, err = str.Seek(-12, io.SeekEnd)
	if err != nil {
		log.Fatalf("Seek error: %s", err)
	}
	fmt.Printf("- Moved offset 12 bytes before end: %d\n", newOffset)

	buf = make([]byte, 12)
	n, err = str.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatalf("Read error: %s", err)
	}
	fmt.Printf("- Final read (%d bytes): %q\n", n, buf[:n])
}

func discardExample() {
	str1 := strings.NewReader(strings.Repeat("a", 42))

	// Discard: count bytes cheaply without storing output
	n, _ := io.Copy(io.Discard, str1)
	fmt.Printf("- Discarded %d bytes\n", n)
}

func extraExample() {
	fmt.Println("\nExtra example:")

	fmt.Println("\nCopy:")
	copyExample()

	fmt.Println("\nCopyN (8 bytes):")
	copyNExample()

	fmt.Println("\nSeek:")
	seekExample()

	fmt.Println("\nDiscard:")
	discardExample()
}
