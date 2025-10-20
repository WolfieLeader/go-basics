package format

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type BookV1 struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Authors     []string  `json:"authors"`
	PublishedAt time.Time `json:"publishedAt"`
	Genres      []string  `json:"genres"`
	IsOnline    bool      `json:"isOnline"`
}

const BOOK_V1_FILE = "book_v1.json"

func JsonV1EncoderExample() {
	book := BookV1{
		Id:          "b1c2d3e4-f5a6-7890-b1c2-d3e4f5a67890",
		Title:       "The Go Programming Language",
		Authors:     []string{"Alan Donovan", "Brian Kernighan"},
		PublishedAt: time.Date(2015, time.November, 20, 0, 0, 0, 0, time.UTC),
		Genres:      []string{"Programming", "Technology"},
		IsOnline:    true,
	}
	fmt.Printf("- Book Struct: %+v\n\n", book)

	file := writerToFile(BOOK_V1_FILE)
	defer file.Close()

	fmt.Println("- JSON:")
	encoder := json.NewEncoder(io.MultiWriter(file, os.Stdout))
	encoder.SetIndent("", "  ")  // Similar to MarshalIndent, can be omitted for compact
	encoder.SetEscapeHTML(false) // Disable HTML escaping

	if err := encoder.Encode(book); err != nil {
		log.Fatalf("JSON Encode error: %v", err)
	}
	fmt.Printf("- JSON written to file %s using Encoder\n", BOOK_V1_FILE)
}

func JsonV1DecoderExample() {
	file := readerFromFile(BOOK_V1_FILE)
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields() // Strict mode for unknown fields

	var book BookV1
	if err := decoder.Decode(&book); err != nil {
		log.Fatalf("JSON Decode error: %v", err)
	}

	if decoder.More() { // Check for multiple JSON objects
		log.Fatalf("JSON Decode error: multiple JSON objects in file")
	}

	fmt.Printf("- Book Struct from JSON file %s using Decoder: %+v\n", BOOK_V1_FILE, book)
}
