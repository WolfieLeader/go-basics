package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"
)

func ioPipeExample() {
	fmt.Println("\nio.Pipe (producer/consumer):")
	r, w := io.Pipe()
	defer r.Close()

	go func() {
		defer w.Close()
		for i := range 3 {
			fmt.Fprintf(w, "- %d", i+1)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	buf := make([]byte, 16)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("- Read %d bytes: %q\n", n, buf[:n])
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("Read error: %s", err)
		}
	}
}
