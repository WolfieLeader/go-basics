package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func jsonWriteExample() {
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

	writeToFile(USER1_FILE, pretty)
}

func jsonReadExample() {
	data := readFromFile(USER1_FILE)

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}
	fmt.Printf("- User Struct from JSON: %+v\n", user)
}
