package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func jsonEncoderExample() {
	date := time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC)
	user := User{
		Name:       "Jane Doe",
		Email:      "jane.doe@example.com",
		Role:       RoleSF,
		Age:        27,
		IsVerified: false,
		Contact:    Contact{Phone: "+1-555-5678", Address: "123 Main St, Anytown, USA"},
		Games:      map[string]int{"vs Miami": 20, "vs Boston": 25},
		CreatedAt:  time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:  &date,
		Password:   "anothersecret",
	}
	fmt.Printf("- User Struct: %+v\n\n", user)

	file := fileWriter(USER2_FILE)
	defer file.Close()

	fmt.Println("- JSON:")
	encoder := json.NewEncoder(io.MultiWriter(file, os.Stdout))
	encoder.SetIndent("", "  ")  // Similar to MarshalIndent, can be omitted for compact
	encoder.SetEscapeHTML(false) // Disable HTML escaping

	if err := encoder.Encode(user); err != nil {
		log.Fatalf("JSON Encode error: %v", err)
	}
	fmt.Printf("- JSON written to file %s using Encoder\n", USER2_FILE)
}

func jsonDecoderExample() {
	file := fileReader(USER2_FILE)
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields() // Strict mode for unknown fields

	var user User
	if err := decoder.Decode(&user); err != nil {
		log.Fatalf("JSON Decode error: %v", err)
	}

	// You can also check if there is more data with decoder.More() but to keep it simple, we assume single object here

	fmt.Printf("- User Struct from JSON file %s using Decoder: %+v\n", USER2_FILE, user)
}
