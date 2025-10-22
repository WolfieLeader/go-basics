package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	dirPath      = "files"
	USER1_FILE   = "user1.json"
	USER2_FILE   = "user2.json"
	BOOKS_FILE   = "books.csv"
	CATALOG_FILE = "catalog.xml"
)

func writeToFile(fileName string, data []byte) {
	file, err := os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		log.Fatalf("Create file error: %v", err)
	}
	defer file.Close()

	n, err := file.Write(data)
	if err != nil {
		log.Fatalf("Write file error: %v", err)
	}
	fmt.Printf("- Wrote %d bytes to file %s\n", n, fileName)
}

func fileWriter(fileName string) io.WriteCloser {
	file, err := os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		log.Fatalf("Create file error: %v", err)
	}
	return file
}

func readFromFile(fileName string) []byte {
	data, err := os.ReadFile(filepath.Join(dirPath, fileName))
	if err != nil {
		log.Fatalf("Read file error: %v", err)
	}
	return data
}

func fileReader(fileName string) io.ReadCloser {
	file, err := os.Open(filepath.Join(dirPath, fileName))
	if err != nil {
		log.Fatalf("Open file error: %v", err)
	}
	return file
}

func main() {
	// JSON v2 examples will be added in 1.26
	fmt.Println("JSON Write Example:")
	jsonWriteExample()
	fmt.Println()

	fmt.Println("JSON Read Example:")
	jsonReadExample()
	fmt.Println()

	fmt.Println("JSON Encoder Write Example:")
	jsonEncoderExample()
	fmt.Println()

	fmt.Println("JSON Decoder Read Example:")
	jsonDecoderExample()
	fmt.Println()

	fmt.Println("CSV Write Example:")
	csvWriteExample()
	fmt.Println()

	fmt.Println("CSV Read Example:")
	csvReadExample()
	fmt.Println()

	fmt.Println("XML Write Example:")
	xmlWriteExample()
	fmt.Println()

	fmt.Println("XML Read Example:")
	xmlReadExample()
	fmt.Println()
}
