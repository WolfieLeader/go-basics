package sorting

import (
	"cmp"
)

type arrayType[T cmp.Ordered] []T

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integers interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Numbers interface {
	Integers | Float
}

func copyArray[T cmp.Ordered](arr []T) (arrayType[T], int) {
	length := len(arr)
	out := make(arrayType[T], length)
	copy(out, arr)
	return out, length
}

func (a *arrayType[T]) swap(index1, index2 int) {
	(*a)[index1], (*a)[index2] = (*a)[index2], (*a)[index1]
}

func (a arrayType[T]) findMinMax() (min T, max T) {
	min, max = a[0], a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
