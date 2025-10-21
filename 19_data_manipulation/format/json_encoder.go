package format

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Book struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Authors     []string  `json:"authors"`
	PublishedAt time.Time `json:"publishedAt"`
	Genres      []string  `json:"genres"`
	IsOnline    bool      `json:"isOnline"`
}

const BOOK_FILE = "book.json"

func JsonV1EncoderExample() {
	book := Book{
		Id:          "b1c2d3e4-f5a6-7890-b1c2-d3e4f5a67890",
		Title:       "The Go Programming Language",
		Authors:     []string{"Alan Donovan", "Brian Kernighan"},
		PublishedAt: time.Date(2015, time.November, 20, 0, 0, 0, 0, time.UTC),
		Genres:      []string{"Programming", "Technology"},
		IsOnline:    true,
	}
	fmt.Printf("- Book Struct: %+v\n\n", book)

	file := writerToFile(BOOK_FILE)
	defer file.Close()

	fmt.Println("- JSON:")
	encoder := json.NewEncoder(io.MultiWriter(file, os.Stdout))
	encoder.SetIndent("", "  ")  // Similar to MarshalIndent, can be omitted for compact
	encoder.SetEscapeHTML(false) // Disable HTML escaping

	if err := encoder.Encode(book); err != nil {
		log.Fatalf("JSON Encode error: %v", err)
	}
	fmt.Printf("- JSON written to file %s using Encoder\n", BOOK_FILE)
}

func JsonV1DecoderExample() {
	file := readerFromFile(BOOK_FILE)
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields() // Strict mode for unknown fields

	var book Book
	if err := decoder.Decode(&book); err != nil {
		log.Fatalf("JSON Decode error: %v", err)
	}

	// You can also check if there is more data with decoder.More() but to keep it simple, we assume single object here

	fmt.Printf("- Book Struct from JSON file %s using Decoder: %+v\n", BOOK_FILE, book)
}
