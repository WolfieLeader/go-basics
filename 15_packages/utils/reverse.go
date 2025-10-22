package utils

import "fmt"

func Reverse(s string) string {
	r := []rune(s)
	for start, end := 0, len(r)-1; start < end; start, end = start+1, end-1 {
		r[start], r[end] = r[end], r[start]
	}
	return string(r)
}

func privateFunction() string {
	return "This is a private function"
}

// Init function executes when the package is imported
// It is often used for setup tasks.
func init() {
	fmt.Println("- Init function in utils package 🧪")
}
