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
	out := make([]int, 0)
	for v := range fibGenerator(23) {
		out = append(out, v)
	}
	fmt.Printf("- %v\n", out)
}
