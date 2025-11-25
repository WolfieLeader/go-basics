package patterns

import "fmt"

func fibGenerator(count int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		if count <= 0 {
			return
		}

		a, b := 0, 1
		for range count {
			ch <- a
			a, b = b, a+b
		}
	}()

	return ch
}

func GeneratorExample() {
	ch := make([]int, 0)
	for v := range fibGenerator(23) {
		ch = append(ch, v)
	}
	fmt.Printf("- Fibonacci numbers: %v\n", ch)
}
