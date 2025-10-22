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

func multiReaderExample() {
	var str1 = strings.NewReader("Foo ")
	var str2 = strings.NewReader("Bar ")
	var str3 = strings.NewReader("Baz")
	mr := io.MultiReader(str1, str2, str3)
	read(mr)
}

func readAllExample() {
	str := strings.NewReader("Hello World! I'm using a Reader!")
	// ReadAll reads from r until EOF and returns the data it read.
	// If you encounter large streams, consider using a buffer instead to avoid high memory usage.
	data, err := io.ReadAll(str)
	if err != nil {
		log.Fatalf("ReadAll error: %s", err)
	}
	fmt.Printf("- ReadAll data: %q\n", data)
}

func teeReaderExample() {
	str := strings.NewReader("- Hello os.Stdout, I'm using a TeeReader!\n")
	tr := io.TeeReader(str, os.Stdout)
	if _, err := io.ReadAll(tr); err != nil {
		log.Fatalf("TeeReader ReadAll error: %s", err)
	}
}

func limitReaderExample() {
	str := strings.NewReader("LimitReader Example: Lorem ipsum dolor sit amet.")
	lr := io.LimitReader(str, 17) // Reads till EOF or 17 bytes, whichever comes first
	read(lr)
}

func readAtExample() {
	str := strings.NewReader("I'm using a Reader!")
	buf := make([]byte, 6)
	n, err := str.ReadAt(buf, 10)
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("ReadAt error: %s", err)
	}
	fmt.Printf("- ReadAt %d bytes: %q\n", n, buf[:n])
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

	fmt.Println("io.MultiReader Read:")
	multiReaderExample()
	fmt.Println()

	fmt.Println("io.ReadAll Read:")
	readAllExample()
	fmt.Println()

	fmt.Println("io.TeeReader Read:")
	teeReaderExample()
	fmt.Println()

	fmt.Println("io.LimitReader Read:")
	limitReaderExample()
	fmt.Println()

	fmt.Println("io.ReaderAt Read:")
	readAtExample()
	fmt.Println()
}
