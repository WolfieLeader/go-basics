package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
)

func zLibExample() {
	data := []byte(LOREM_IPSUM)
	fmt.Printf("- Original Text Size: %d bytes\n", len(data))

	var buf bytes.Buffer
	zlibWriter := zlib.NewWriter(&buf) // uses `zlib.DefaultCompression`

	if _, err := zlibWriter.Write(data); err != nil {
		log.Fatalf("Zlib write error: %v", err)
	}

	if err := zlibWriter.Close(); err != nil {
		log.Fatalf("Zlib close error: %v", err)
	}

	compressed := buf.Bytes()
	fmt.Printf("- Compressed Size: %d bytes\n", len(compressed))

	newBuf := bytes.NewReader(compressed)
	zlibReader, err := zlib.NewReader(newBuf) // Can also use `&buf`, buf for safety is better
	if err != nil {
		log.Fatalf("Zlib reader error: %v", err)
	}
	defer zlibReader.Close()

	decompressed, err := io.ReadAll(zlibReader)
	if err != nil {
		log.Fatalf("Zlib read error: %v", err)
	}

	fmt.Printf("- Decompressed matches original? %v\n", bytes.Equal(decompressed, data))
}
