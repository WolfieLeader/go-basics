package main

import "fmt"

// Define “ordered” types (numbers & strings)
type Ordered interface {
    ~int | ~float64 | ~string
}

// Generic min function for anything ordered
func Min[T Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Min(3, 7))        // 3
    fmt.Println(Min(3.14, 2.71))  // 2.71
    fmt.Println(Min("go", "go!"))  // "go"
}
