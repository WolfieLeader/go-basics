package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func ifExample() {
	age := 20
	if age >= 18 {
		fmt.Println("- You are an adult.")
	} else {
		fmt.Println("- You are a minor.")
	}
}

func switchExample() {
	switch runtime.GOOS {
	case "darwin":
		fmt.Println("- You are running on macOS")
	case "linux":
		fmt.Println("- You are running on Linux")
	case "windows":
		fmt.Println("- You are running on Windows")
	default: // If no case matches, the default case is executed
		fmt.Println("- You are running on an unknown OS")
	}
}

func ifWithStatementExample() {
	if num, err := strconv.Atoi("123"); err == nil {
		fmt.Printf("- Converted string to int: %d\n", num)
	} else {
		fmt.Println("- Error converting string to int:", err)
	}
}

func switchTrueExample() {
	score := 750
	switch {
	case score <= 0: // 0 or below
		fmt.Println("- Level: New")
	case score <= 100: // 1 - 100
		fmt.Println("- Level: Bronze")
	case score <= 500: // 101 - 500
		fmt.Println("- Level: Silver")
	case score <= 1000: // 501 - 1000
		fmt.Println("- Level: Gold")
	case score <= 5000: // 1001 - 5000
		fmt.Println("- Level: Diamond")
	}
}

func main() {
	fmt.Println("If Example:")
	ifExample()
	fmt.Println()

	fmt.Println("Switch Example:")
	switchExample()
	fmt.Println()

	fmt.Println("If with Statement Example:")
	ifWithStatementExample()
	fmt.Println()

	fmt.Println("Switch True Example:")
	switchTrueExample()
	fmt.Println()
}
