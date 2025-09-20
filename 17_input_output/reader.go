package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Reader is an interface with one method: Read(p []byte) (n int, err error)
// It is implemented by many types, including os.File, os.Stdin, bytes.Buffer, strings.Reader, http.Request.Body etc.
func read(r io.Reader) {
	// Create a buffer to hold the read bytes
	// Here we limit it to 16 bytes for demonstration
	buf := make([]byte, 16)

	//Common pattern to read from a Reader until EOF (end of file)
	for {
		n, err := r.Read(buf)
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
	fmt.Println("Finished reading.")
}

func osFileReaderExample() {
	file, err := os.Open("texts/lorem-ipsum.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	defer func() {
		// type *os.File implements Closer as well
		if err := file.Close(); err != nil {
			log.Fatalf("Failed to close file: %s", err)
		}
	}()

	read(file)
}

func stringsReaderExample() {
	str := strings.NewReader("Hello World! I'm using a Reader!")
	read(str)
}

func httpResponseBodyExample() {
	resp, err := http.Get("https://api.github.com/zen")
	if err != nil {
		log.Fatalf("HTTP GET error: %s", err)
	}
	defer resp.Body.Close()

	read(resp.Body)
}

func readerExample() {
	fmt.Println("\nReader example:")

	fmt.Println("\nos.File Read:")
	osFileReaderExample()

	fmt.Println("\nstrings.Reader Read:")
	stringsReaderExample()

	fmt.Println("\nhttp.Response.Body Read:")
	httpResponseBodyExample()
}
