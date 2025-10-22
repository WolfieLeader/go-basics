package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Role string
}

var users = []User{{Name: "Alice", Role: "user"}, {Name: "Bob", Role: "admin"}}

var ErrorNotFound = errors.New("not found")

type ErrorUnauthorized struct {
	Role string
}

func (e *ErrorUnauthorized) Error() string {
	return fmt.Sprintf("Unauthorized access for role: %s", e.Role)
}

func watchInfo(userName string) (string, error) {
	for _, user := range users {
		if user.Name == userName {
			return fmt.Sprintf("User %s has role %s", user.Name, user.Role), nil
		}
	}
	// Error wrapping is used to provide more context about the error without losing the original error
	return "", fmt.Errorf("%s user is %w", userName, ErrorNotFound)
}

func watchAdminInfo(userName string) (string, error) {
	for _, user := range users {
		if user.Name == userName {
			if user.Role != "admin" {
				return "", &ErrorUnauthorized{Role: user.Role}
			}
			return fmt.Sprintf("Admin user %s has full access", user.Name), nil
		}
	}
	return "", ErrorNotFound
}

func customErrorExample() {
	charlie, err := watchInfo("Ghost")

	if errors.Is(err, ErrorNotFound) { // Check if error is exactly ErrorNotFound
		fmt.Printf("- Not found wrapped error: %s\n", err)
		fmt.Printf("- Original not found error: %v\n", errors.Unwrap(err))
	} else if err != nil { // Check for other errors
		fmt.Println("- Error occurred:", err)
	} else { // No error
		fmt.Println("- Found user:", charlie)
	}

	alice, err := watchAdminInfo("Alice")
	var errorUnauthorized *ErrorUnauthorized

	if errors.As(err, &errorUnauthorized) { // Check and extract if error is of type *ErrorUnauthorized
		fmt.Println("- Unauthorized access:", errorUnauthorized.Error())
	} else if err != nil { // Other errors
		fmt.Println("- Error occurred:", err)
	} else { // No error
		fmt.Println("- Admin info for Alice:", alice)
	}

	bob, err := watchAdminInfo("Bob")
	if err != nil {
		fmt.Println("- Error occurred:", err)
	} else {
		fmt.Println("- Admin info for Bob:", bob)
	}
}
