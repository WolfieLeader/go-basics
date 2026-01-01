package sorting

import "cmp"

func Cocktail[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	left, right := 0, length-1
	for left < right {
		swapped := false
		lastRight := left

		// left -> right
		for i := left; i < right; i++ {
			if array[i] > array[i+1] {
				array.swap(i+1, i)
				swapped = true
				lastRight = i
			}
		}
		if !swapped {
			break
		}
		right = lastRight // everything after lastRight is in place

		// right -> left
		swapped = false
		lastLeft := right
		for j := right; j > left; j-- {
			if array[j] < array[j-1] {
				array.swap(j, j-1)
				swapped = true
				lastLeft = j
			}
		}
		if !swapped {
			break
		}
		left = lastLeft // everything before lastLeft is in place
	}
	return array
}
