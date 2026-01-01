package searching

import "cmp"

func LinearSearch[T cmp.Ordered](array []T, value T) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}
