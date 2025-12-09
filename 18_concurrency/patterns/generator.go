package patterns

import "fmt"

func generator(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, v := range nums {
			ch <- v
		}
	}()

	return ch
}

func GeneratorExample() {
	nums := make([]int, 0)
	for v := range generator(1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144) {
		nums = append(nums, v)
	}
	fmt.Printf("- Fibonacci numbers: %v\n", nums)
}
