package compress

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
)

func GzipExample() {
	data := []byte(LOREM_IPSUM)
	fmt.Printf("- Original Text Size: %d bytes\n", len(data))

	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf) // uses `gzip.DefaultCompression`

	// optional: set header fields for extra metadata
	gzipWriter.Name = "lorem_ipsum.txt"
	gzipWriter.Comment = "Gzip Example"

	if _, err := gzipWriter.Write(data); err != nil {
		log.Fatalf("Gzip write error: %v", err)
	}

	if err := gzipWriter.Close(); err != nil {
		log.Fatalf("Gzip close error: %v", err)
	}

	compressed := buf.Bytes()
	fmt.Printf("- Compressed Size: %d bytes\n", len(compressed))

	newBuf := bytes.NewReader(compressed)
	gzipReader, err := gzip.NewReader(newBuf) // Can also use `&buf`, buf for safety is better
	if err != nil {
		log.Fatalf("Gzip reader error: %v", err)
	}
	defer gzipReader.Close()

	fmt.Printf("- Gzip Header Name: %q\n", gzipReader.Name)
	fmt.Printf("- Gzip Header Comment: %q\n", gzipReader.Comment)

	decompressed, err := io.ReadAll(gzipReader)
	if err != nil {
		log.Fatalf("Gzip read error: %v", err)
	}

	fmt.Printf("- Decompressed matches original? %v\n", bytes.Equal(decompressed, data))
}
