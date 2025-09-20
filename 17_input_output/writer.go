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
func write(w io.Writer, data string) {
	n, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalf("Write error: %s", err)
	}
	fmt.Printf("- Wrote %d bytes\n", n)
}

func osFileWriterExample() {
	file, err := os.Create("texts/out.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	// type *os.File implements Closer as well
	defer file.Close()

	write(file, "Go was designed at Google, ")
	write(file, "by Robert Griesemer, Rob Pike, and Ken Thompson.")
}

func stringsBuilderExample() {
	var sb strings.Builder
	write(&sb, "Hello ")
	write(&sb, "World!")
	sb.WriteString(" Added with WriteString method.")
	fmt.Printf("Final string: %q\n", sb.String())
}

func osStdoutExample() {
	write(os.Stdout, "- Hey ")
	write(os.Stdout, "- Console ")
}

func multiWriterExample() {
	var sb1, sb2, sb3 strings.Builder
	mw := io.MultiWriter(&sb1, &sb2, &sb3)

	for i := range 3 {
		write(mw, fmt.Sprintf(" %d", i+1))
	}
	write(mw, " ")

	fmt.Printf("Final strings:\n")
	for i, sb := range []*strings.Builder{&sb1, &sb2, &sb3} {
		fmt.Printf("- sb%d: %q\n", i+1, sb.String())
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

	fmt.Println("\nMultiWriter Write:")
	multiWriterExample()
}
