package sorting

import "cmp"

func Gnome[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	i := 1
	for i < length {
		if i == 0 || array[i] >= array[i-1] {
			i++
		} else {
			array.swap(i, i-1)
			i--
		}
	}

	return array
}
