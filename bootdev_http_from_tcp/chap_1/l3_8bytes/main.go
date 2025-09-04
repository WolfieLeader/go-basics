package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	buf := make([]byte, 8)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			fmt.Printf("read: %s\n", buf[:n])
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
	}
}
