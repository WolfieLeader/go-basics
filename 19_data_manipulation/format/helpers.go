package format

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const dirPath = "./format/files"

func writeBytesToFile(fileName string, data []byte) {
	f, err := os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		log.Fatalf("Create file error: %v", err)
	}
	defer f.Close()

	n, err := f.Write(data)
	if err != nil {
		log.Fatalf("Write file error: %v", err)
	}
	fmt.Printf("- Wrote %d bytes to file %s\n", n, fileName)
}

func writerToFile(fileName string) io.WriteCloser {
	f, err := os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		log.Fatalf("Create file error: %v", err)
	}
	return f
}

func readBytesFromFile(fileName string) []byte {
	data, err := os.ReadFile(filepath.Join(dirPath, fileName))
	if err != nil {
		log.Fatalf("Read file error: %v", err)
	}
	return data
}

func readerFromFile(fileName string) io.ReadCloser {
	f, err := os.Open(filepath.Join(dirPath, fileName))
	if err != nil {
		log.Fatalf("Open file error: %v", err)
	}
	return f
}
