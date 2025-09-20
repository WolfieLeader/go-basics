package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Writer is an interface with one method: Write(p []byte) (n int, err error)
// It is implemented by many types, including *os.File, strings.Builder, http.ResponseWriter, bytes.Buffer, etc.
// if n < len(p), err MUST be non-nil
func write(w io.Writer, data string) error {
	n, err := w.Write([]byte(data))
	if err != nil {
		return err
	}
	fmt.Printf("- Wrote %d bytes\n", n)
	return nil
}

func osFileWriterExample() {
	file, err := os.Create("texts/out.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Failed to close file: %s", err)
		}
	}()

	if err := write(file, "Go was designed at Google, "); err != nil {
		log.Fatalf("Write error: %s", err)
	}

	if err := write(file, "by Robert Griesemer, Rob Pike, and Ken Thompson."); err != nil {
		log.Fatalf("Write error: %s", err)
	}
}

func stringsBuilderExample() {
	var sb strings.Builder
	if err := write(&sb, "Hello "); err != nil {
		log.Fatalf("Write error: %s", err)
	}
	if err := write(&sb, "World!"); err != nil {
		log.Fatalf("Write error: %s", err)
	}
	fmt.Printf("Final string: %q\n", sb.String())
}

func osStdoutExample() {
	if err := write(os.Stdout, "- Hey "); err != nil {
		log.Fatalf("Write error: %s", err)
	}

	if err := write(os.Stdout, "- Console "); err != nil {
		log.Fatalf("Write error: %s", err)
	}
}

func writerExample() {
	fmt.Println("\nWriter Example:")

	fmt.Println("\nos.File Write:")
	osFileWriterExample()

	fmt.Println("\nstrings.Builder Write:")
	stringsBuilderExample()

	fmt.Println("\nos.Stdout Write:")
	osStdoutExample()
}
