package sorting

import "cmp"

func Insertion[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	for pass := 1; pass < length; pass++ {
		value := array[pass]
		i := pass - 1

		for i >= 0 && array[i] > value {
			array[i+1] = array[i]
			i--
		}

		array[i+1] = value
	}
	return array
}
