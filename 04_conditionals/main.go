package main

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
)

// Common pattern for error handling in Go
func divide(a, b int) (float64, error) {
	// No parentheses() are needed
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return float64(a) / float64(b), nil
}

func convertToInt(str string) (int, error) {
	// You can declare and initialize a var inside an if statement, it's good for short-lived variables
	if num, err := strconv.Atoi(str); err == nil {
		return num, nil
	} else {
		return 0, err
	}
}

func printOS() {
	// Switch statements are like if statements
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("You are running on macOS")
	case "linux":
		fmt.Println("You are running on Linux")
	case "windows":
		fmt.Println("You are running on Windows")
	default: // If no case matches, the default case is executed
		fmt.Println("You are running on an unknown OS: ", os)
	}
}

// switch without an expression is like a series of if statements.
func levelByScore(x int) string {
	switch {
	case x <= 0:
		return "New"
	case x <= 100:
		return "Bronze"
	case x <= 500:
		return "Silver"
	case x <= 1000:
		return "Gold"
	case x <= 5000:
		return "Platinum"
	case x <= 10000:
		return "Diamond"
	default:
		return "Master"
	}
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		return
	}
	fmt.Printf("10 divided by 2 is: %.2f\n", result)

	num, err := convertToInt("42")
	if err != nil {
		return
	}
	fmt.Printf("Converted string to int: %d\n", num)

	printOS()
	fmt.Println("Your level based on score 1500 is:", levelByScore(1500))
	fmt.Println("Your level based on score 5000 is:", levelByScore(5000))
}
