package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
)

func flateExample() {
	data := []byte(LOREM_IPSUM)
	fmt.Printf("- Original Text Size: %d bytes\n", len(data))

	var buf bytes.Buffer
	flateWriter, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		log.Fatalf("Flate writer error: %v", err)
	}

	if _, err := flateWriter.Write(data); err != nil {
		log.Fatalf("Flate write error: %v", err)
	}

	if err := flateWriter.Close(); err != nil {
		log.Fatalf("Flate close error: %v", err)
	}

	compressed := buf.Bytes()
	fmt.Printf("- Compressed Size: %d bytes\n", len(compressed))

	newBuf := bytes.NewReader(compressed)
	flateReader := flate.NewReader(newBuf) // Can also use `&buf`, buf for safety is better
	defer flateReader.Close()

	decompressed, err := io.ReadAll(flateReader)
	if err != nil {
		log.Fatalf("Flate read error: %v", err)
	}

	fmt.Printf("- Decompressed matches original? %v\n", bytes.Equal(decompressed, data))
}
