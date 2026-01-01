package sorting

import "cmp"

// ------ Lomuto Partition -----

func Quick[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	lomutoQuickSort(array, 0, length-1)
	return array
}

func lomutoQuickSort[T cmp.Ordered](array arrayType[T], low int, high int) {
	if low >= high {
		return
	}
	pivot := lomutoPartition(array, low, high)
	lomutoQuickSort[T](array, low, pivot-1)
	lomutoQuickSort[T](array, pivot+1, high)
}

func lomutoPartition[T cmp.Ordered](array arrayType[T], low int, high int) int {
	median := medianOfThree(array, low, high)
	array.swap(median, high)
	pivotVal := array[high]

	i := low
	for j := low; j < high; j++ {
		if array[j] < pivotVal {
			array.swap(i, j)
			i++
		}
	}

	array.swap(i, high)
	return i
}

// ------ Hoare Partition ------

func QuickHoare[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	hoareQuickSort(array, 0, length-1)
	return array
}

func hoareQuickSort[T cmp.Ordered](array arrayType[T], low int, high int) {
	if low >= high {
		return
	}

	pivot := hoarePartition(array, low, high)
	hoareQuickSort[T](array, low, pivot)
	hoareQuickSort[T](array, pivot+1, high)
}

func hoarePartition[T cmp.Ordered](array arrayType[T], low int, high int) int {
	median := medianOfThree(array, low, high)
	pivotVal := array[median]

	left, right := low, high
	for {
		for array[left] < pivotVal {
			left++
		}
		for array[right] > pivotVal {
			right--
		}
		if left >= right {
			return right
		}
		array.swap(left, right)
		left++
		right--
	}
}

// ------ 3 Way (Dutch Flag) -----

func Quick3Way[T cmp.Ordered](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	threeWayQuickSort(array, 0, length-1)
	return array
}

func threeWayQuickSort[T cmp.Ordered](array arrayType[T], low int, high int) {
	if low >= high {
		return
	}

	lesser, greater := threeWayPartition(array, low, high)
	threeWayQuickSort[T](array, low, lesser-1)
	threeWayQuickSort[T](array, greater+1, high)
}

func threeWayPartition[T cmp.Ordered](array arrayType[T], low int, high int) (int, int) {
	median := medianOfThree(array, low, high)
	pivotVal := array[median]

	lesser, equal, greater := low, low, high
	for equal <= greater {
		switch {
		case array[equal] < pivotVal:
			array.swap(lesser, equal)
			lesser++
			equal++
		case array[equal] > pivotVal:
			array.swap(equal, greater)
			greater--
		default:
			equal++
		}
	}

	return lesser, greater
}

// ------ Shared Utilities ------

func medianOfThree[T cmp.Ordered](array arrayType[T], low, high int) int {
	mid := low + (high-low)/2

	if array[mid] < array[low] {
		low, mid = mid, low
	}
	if array[high] < array[low] {
		low, high = high, low
	}
	if array[high] < array[mid] {
		mid, high = high, mid
	}

	_, _ = low, high // remove warnings
	return mid       // returns the median index without swapping
}
