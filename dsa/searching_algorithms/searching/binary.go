package searching

import "cmp"

func BinarySearch[T cmp.Ordered](array []T, value T) int {
	low, high := 0, len(array)-1

	for low <= high {
		mid := low + (high-low)/2

		if array[mid] == value {
			return mid
		}

		if array[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
