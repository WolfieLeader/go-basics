package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Role string
}

var ErrNotFound = errors.New("User not found")

type UnauthorizedError struct {
	Role string
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("Unauthorized access for role: %s", e.Role)
}

var users = []User{
	{Name: "Alice", Role: "user"},
	{Name: "Bob", Role: "admin"},
}

func watchInfo(userName string) (string, error) {
	for _, user := range users {
		if user.Name == userName {
			return fmt.Sprintf("User %s has role %s", user.Name, user.Role), nil
		}
	}
	return "", ErrNotFound
}

func watchAdminInfo(userName string) (string, error) {
	for _, user := range users {
		if user.Name == userName {
			if user.Role != "admin" {
				return "", &UnauthorizedError{Role: user.Role}
			}
			return fmt.Sprintf("Admin user %s has full access", user.Name), nil
		}
	}
	return "", ErrNotFound
}

func customErrorExample() {
	fmt.Println("\nCustom Error Example:")
	charlie, err := watchInfo("Charlie")
	// errors.Is is used to check if the error is exactly ErrNotFound by value comparison
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("Not found error: %s\n", err)
	// Here we check if there is another type of error
	} else if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Found user:", charlie)
	}

	alice, err := watchAdminInfo("Alice")
	var unauthorizedErr *UnauthorizedError
	// errors.As is used to check if the error is of type UnauthorizedError and to extract it to unauthorizedErr
	if errors.As(err, &unauthorizedErr) {
		fmt.Println("Unauthorized access:", unauthorizedErr.Error())
	// Here we check if there is another type of error
	} else if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Admin info for Alice:", alice)
	}

	bob, err := watchAdminInfo("Bob")
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Admin info for Bob:", bob)
	}
}