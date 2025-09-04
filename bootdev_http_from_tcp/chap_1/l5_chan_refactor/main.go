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

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
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
					out <- sb.String()
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
			out <- sb.String()
		}
	}()
	return out
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	for line := range getLinesChannel(f) {
		fmt.Printf("read: %s\n", line)
	}
}
