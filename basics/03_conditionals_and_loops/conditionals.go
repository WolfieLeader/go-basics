package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func absolute(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// This is common return pattern in Go where you return a value and an error.
func convertToInt(str string) (int, error) {
	// In Go you can declare and initialize a var inside an if statement this is good for short-lived variables
	if num, err := strconv.Atoi(str); err == nil {
		return num, nil
	} else {
		return 0, err
	}
}

func printOS() {
	// Switch statements are like if statements with multiple conditions.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("You are running on macOS")
	case "linux":
		fmt.Println("You are running on Linux")
	case "windows":
		fmt.Println("You are running on Windows")
	default: // Default case is executed if none of the above cases match
		fmt.Println("You are running on an unknown OS: ", os)
	}
}
