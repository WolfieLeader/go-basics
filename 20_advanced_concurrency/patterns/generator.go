package patterns

import "fmt"

// Generator is a function that returns a channel
func evenNumbersGenerator(start, count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range count {
			out <- start + i*2
		}
	}()

	return out //Returns a read-only channel
}

func GeneratorExample() {
	for v := range evenNumbersGenerator(10, 5) {
		fmt.Println("-", v)
	}
}
