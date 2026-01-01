package sorting

import "cmp"

func Bubble[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	for pass := 0; pass < length-1; pass++ {
		swapped := false

		for i := 0; i < length-1-pass; i++ {
			if array[i] > array[i+1] {
				array.swap(i+1, i)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
	return array
}
