package format

import (
	"encoding/json"
	"fmt"
	"log"
)

type UserV1 struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	Age         int               `json:"age"`
	IsAdmin     bool              `json:"isAdmin"`
	Interests   []string          `json:"interests"`
	Preferences PreferencesV1     `json:"preferences"`
	Network     map[string]string `json:"network"`
	Password    string            `json:"-"`                  // Exclude from JSON
	IsActive    string            `json:"isActive,omitempty"` // Omit if empty when marshaling
}

type PreferencesV1 struct {
	Notifications bool   `json:"notifications"`
	Theme         string `json:"theme"`
}

const USER_V1_FILE = "user_v1.json"

func JsonV1WriteExample() {
	var user = UserV1{
		Id:          "fad07fb3-784c-406c-adb9-b9e765e5b380",
		Name:        "John Doe",
		Email:       "john.doe@example.com",
		Age:         30,
		IsAdmin:     false,
		Interests:   []string{"tech", "podcast", "beautiful models"},
		Preferences: PreferencesV1{Notifications: true, Theme: "dark"},
		Network: map[string]string{
			"ce10b4dd-384f-4644-b084-42f5c5efe045": "friend",
			"a4f5c6d7-e8f9-1011-1213-141516171819": "family",
			"1a2b3c4d-5e6f-7081-9201-112131415161": "pet",
		},
		Password: "SuperSecretPassword123!",
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

	writeBytesToFile(USER_V1_FILE, pretty)
}

func JsonV1ReadExample() {
	data := readBytesFromFile(USER_V1_FILE)

	var user UserV1
	if err := json.Unmarshal(data, &user); err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}
	fmt.Printf("- User Struct from JSON: %+v\n", user)
}
