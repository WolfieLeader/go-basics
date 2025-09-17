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
func read(r io.Reader) error {
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
			return err
		}
	}
	fmt.Println("Finished reading.")
	return nil
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

	if err := read(file); err != nil {
		log.Fatalf("Read error: %s", err)
	}
}

func stringsReaderExample() {
	str := strings.NewReader("Hello World! I'm using a Reader!")
	if err := read(str); err != nil {
		log.Fatalf("Read error: %s", err)
	}
}

func httpResponseBodyExample() {
	resp, err := http.Get("https://api.github.com/zen")
	if err != nil {
		log.Fatalf("HTTP GET error: %s", err)
	}

	defer func() {
		// type *http.Response.Body implements Closer as well
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("Failed to close response body: %s", err)
		}
	}()

	if err := read(resp.Body); err != nil {
		log.Fatalf("Read error: %s", err)
	}
}

func readerExample() {
	fmt.Println("os.File Read:")
	osFileReaderExample()
	fmt.Println()

	fmt.Println("strings.Reader Read:")
	stringsReaderExample()
	fmt.Println()

	fmt.Println("http.Response.Body Read:")
	httpResponseBodyExample()
	fmt.Println()
}
