package sorting

import "cmp"

func Comb[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	gap := length
	swapped := true

	for gap > 1 || swapped {
		gap = max((gap*10)/13, 1)

		swapped = false
		for i := 0; i+gap < length; i++ {
			if array[i] > array[i+gap] {
				array.swap(i+gap, i)
				swapped = true
			}
		}
	}

	return array
}
