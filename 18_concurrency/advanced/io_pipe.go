package advanced

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"
)

// Synchronize between Readers and Writers
func IoPipeExample() {
	r, w := io.Pipe()
	defer r.Close()

	// Producer
	go func() {
		defer w.Close()
		for i := 1; i <= 512; i *= 2 {
			fmt.Fprintf(w, "%d", i*i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Consumer
	buf := make([]byte, 8)
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
