package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	var sb strings.Builder
	buf := make([]byte, 8)

	for {
		n, err := f.Read(buf)
		if n > 0 {
			chunk := buf[:n]
			parts := bytes.Split(chunk, []byte{'\n'})

			for i := 0; i < len(parts)-1; i++ {
				sb.Write(parts[i])
				fmt.Printf("read: %s\n", sb.String())
				sb.Reset()
			}

			sb.Write(parts[len(parts)-1])
		}

		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
	}

	if sb.Len() > 0 {
		fmt.Printf("read: %s\n", sb.String())
	}
}
