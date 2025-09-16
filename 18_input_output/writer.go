package main

import (
	"fmt"
	"log"
	"os"
)

func writerExample() {
	fmt.Println("\nWriter Example:")
	// Writer is an interface with one method: Write(p []byte) (n int, err error)
	// It is implemented by many types, including *os.File, bytes.Buffer, strings.Builder, etc.
	// Here we use os.Create to get a *os.File which implements Writer
	file, err := os.Create("out.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	// The *os.File implements Writer as well as Closer (Close() error)
	// We use defer to ensure the file is closed when we're done
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Failed to close file: %s", err)
		}
	}()

	// Data to write
	data := []byte("Go was designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.\n")
	written := 0

	for written < len(data) {
		n, err := file.Write(data[written:])
		if err != nil {
			log.Fatalf("Write error: %s", err)
		}
		written += n
	}

	fmt.Printf("Wrote %d bytes to out.txt\n", written)
	fmt.Println("Finished writing to file.")
}
