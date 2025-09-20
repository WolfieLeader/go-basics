package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func bufioReaderExample() {
	file, err := os.Open("texts/story.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// NewReader method uses NewReaderSize with default buffer size which is 4096 bytes
	br := bufio.NewReaderSize(file, 48) // Custom buffer size of 48 bytes

	peek, err := br.Peek(20) // Peek at the next 20 bytes without advancing the reader
	if err != nil {
		log.Fatalf("Peek error: %s", err)
	}
	fmt.Printf("Peek: %s (buffered=%d)\n", peek, br.Buffered())

	discarded, err := br.Discard(5)
	if err != nil {
		log.Fatalf("Discard error: %s", err)
	}
	fmt.Printf("Discarded %d bytes, buffered=%d\n", discarded, br.Buffered())

	for {
		line, isPrefix, err := br.ReadLine()
		if len(line) > 0 {
			fmt.Printf("- %s (isPrefix=%t, buffered=%d)\n", string(line), isPrefix, br.Buffered())
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("Read error: %s", err)
		}
	}
	if br.Buffered() > 0 {
		remaining, _ := br.Peek(br.Buffered())
		fmt.Printf("Remaining buffered data: %s\n", remaining)
	}
	fmt.Println("Finished buffered reading.")
}

func bufioWriterExample() {
	var sb strings.Builder
	bw := bufio.NewWriterSize(&sb, 16) // Custom buffer size of 16 bytes

	if _, err := bw.WriteString("Buffered"); err != nil {
		log.Fatalf("WriteString error: %s", err)
	}
	if _, err := bw.WriteRune('🔥'); err != nil {
		log.Fatalf("WriteRune error: %s", err)
	}

	fmt.Printf("- There is still %d bytes available in the buffer\n", bw.Available())

	// Flush writes to the underlying writer
	if err := bw.Flush(); err != nil {
		log.Fatalf("Flush error: %s", err)
	}
	fmt.Printf("- Output: %s\n", sb.String())

	bw.Reset(os.Stdout) // Reset to write to standard output
	bw.WriteString("- Since it's larger than 16 bytes, it will auto-flush the buffer")
}

func bufioScannerExample() {
	fmt.Println("Enter text (type 'exit' to quit):")
	bs := bufio.NewScanner(os.Stdin)
	for bs.Scan() {
		line := bs.Text()
		if line == "exit" {
			break
		}
		fmt.Printf("You entered: %s\n", line)
	}
	if err := bs.Err(); err != nil {
		log.Fatalf("Scanner error: %s", err)
	}

	fmt.Println("Scanner finished.")

}

func bufioExample() {
	fmt.Println("\nBufio Example:")

	fmt.Println("\nBufio Reader Example:")
	bufioReaderExample()

	fmt.Println("\nBufio Writer Example:")
	bufioWriterExample()

	fmt.Println("\nBufio Scanner Example:")
	bufioScannerExample()
}
