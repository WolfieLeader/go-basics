package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func readerExample() {
	fmt.Println("\nReader example:")
	// Reader is an interface with one method: Read(p []byte) (n int, err error)
	// It is implemented by many types, including *os.File, bytes.Buffer, strings.Reader, etc.
	// Here we use os.Open to get a *os.File which implements Reader
	file, err := os.Open("lorem-ipsum.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	// The *os.File implements Reader as well as Closer (Close() error)
	// We use defer to ensure the file is closed when we're done
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Failed to close file: %s", err)
		}
	}()

	// Create a buffer to hold the read bytes, here we will limit it to 8 bytes but you can adjust as needed
	buf := make([]byte, 8)

	//Common pattern to read from a Reader until EOF (end of file)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			fmt.Printf("- Read %d bytes: %q\n", n, buf[:n])
		}
		if errors.Is(err, io.EOF) {
			break // End of file reached
		}
		if err != nil {
			log.Fatalf("Read error: %s", err)
		}
	}
	fmt.Println("Finished reading file.")
}
