package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer reader.Close()

	for {
		buf := make([]byte, 8)
		n, err := reader.Read(buf)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		fmt.Printf("read: %s\n", buf[:n])
	}
}
