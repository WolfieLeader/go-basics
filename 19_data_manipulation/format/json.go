package format

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Role string

const (
	RolePG Role = "Point Guard"
	RoleSG Role = "Shooting Guard"
	RoleSF Role = "Small Forward"
	RolePF Role = "Power Forward"
	RoleC  Role = "Center"
)

type Contact struct {
	Phone   string `json:"phone"`
	Address string `json:"address,omitempty"` // Omitempty if empty
}

type User struct {
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Role         Role           `json:"role"`
	JerseyNumber *int           `json:"jerseyNumber,omitempty"` // 0 won't be omitted, nil will be omitted
	Age          int            `json:"age,omitempty"`          // 0 will be omitted
	IsVerified   bool           `json:"isVerified"`
	Interests    []string       `json:"interests"`
	Contact      Contact        `json:"contact"`
	Games        map[string]int `json:"games"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    *time.Time     `json:"updatedAt"`
	Password     string         `json:"-"` // Exclude from JSON
}

const (
	JOHN_DOE_FILE = "user_john_doe.json"
	JANE_DOE_FILE = "user_jane_doe.json"
)

func JsonV1WriteExample() {
	jerseyNumber := 0
	var user = User{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		Role:         RoleSG,
		JerseyNumber: &jerseyNumber,
		Age:          28,
		IsVerified:   true,
		Interests:    []string{"basketball", "coding", "beautiful models"},
		Contact:      Contact{Phone: "+1-555-1234"},
		Games:        map[string]int{"vs New York": 12, "vs Los Angeles": 15, "vs Chicago": 35},
		CreatedAt:    time.Date(2025, time.January, 15, 0, 0, 0, 0, time.UTC),
		Password:     "supersecret",
	}
	fmt.Printf("- User Struct: %+v\n\n", user)

	compact, err := json.Marshal(user) // Most used and efficient
	if err != nil {
		log.Fatalf("Marshal error: %v", err)
	}
	fmt.Printf("- JSON (compact): %s\n\n", compact)

	pretty, err := json.MarshalIndent(user, "", "  ") // Human readable
	if err != nil {
		log.Fatalf("Marshal indent error: %v", err)
	}
	fmt.Printf("- JSON (pretty):\n%s\n\n", pretty)

	writeBytesToFile(JOHN_DOE_FILE, pretty)
}

func JsonV1ReadExample() {
	data := readBytesFromFile(JOHN_DOE_FILE)

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}
	fmt.Printf("- User Struct from JSON: %+v\n", user)
}

func JsonV1EncoderExample() {
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

	file := writerToFile(JANE_DOE_FILE)
	defer file.Close()

	fmt.Println("- JSON:")
	encoder := json.NewEncoder(io.MultiWriter(file, os.Stdout))
	encoder.SetIndent("", "  ")  // Similar to MarshalIndent, can be omitted for compact
	encoder.SetEscapeHTML(false) // Disable HTML escaping

	if err := encoder.Encode(user); err != nil {
		log.Fatalf("JSON Encode error: %v", err)
	}
	fmt.Printf("- JSON written to file %s using Encoder\n", JANE_DOE_FILE)
}

func JsonV1DecoderExample() {
	file := readerFromFile(JANE_DOE_FILE)
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields() // Strict mode for unknown fields

	var user User
	if err := decoder.Decode(&user); err != nil {
		log.Fatalf("JSON Decode error: %v", err)
	}

	// You can also check if there is more data with decoder.More() but to keep it simple, we assume single object here

	fmt.Printf("- User Struct from JSON file %s using Decoder: %+v\n", JANE_DOE_FILE, user)
}
