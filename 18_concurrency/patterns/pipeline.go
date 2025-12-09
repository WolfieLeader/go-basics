package patterns

import "fmt"

// Pipeline - Connecting a series of stages
func double(nums <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range nums {
			ch <- v * 2
		}
	}()
	return ch
}

func square(nums <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range nums {
			ch <- v * v
		}
	}()
	return ch
}

func PipelineExample() {
	nums := make([]int, 0)
	for v := range square(double(generator(1, 2, 3, 4, 5))) {
		nums = append(nums, v)
	}
	fmt.Printf("- Numbers after doubling and squaring: %v\n", nums)
}
