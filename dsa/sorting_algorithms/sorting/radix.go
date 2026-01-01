package sorting

func Radix[T Integers](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	negatives := make([]uint64, 0, length)
	positives := make([]uint64, 0, length)

	for _, value := range array {
		if value < 0 {
			negatives = append(negatives, absToUint64(value))
		} else {
			positives = append(positives, uint64(value))
		}
	}

	k := 0

	if len(negatives) > 0 {
		negatives = radixUint(negatives)
		for i := len(negatives) - 1; i >= 0; i-- {
			array[k] = T(-int64(negatives[i]))
			k++
		}
	}

	if len(positives) > 0 {
		positives = radixUint(positives)
		for _, u := range positives {
			array[k] = T(u)
			k++
		}
	}

	return array
}

func radixUint[T Unsigned](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	minVal, maxVal := array.findMinMax()
	if minVal == maxVal {
		return array
	}

	out := make([]T, length)
	var count [10]int

	for exp := uint64(1); uint64(maxVal)/exp > 0; exp *= 10 {
		for i := range count {
			count[i] = 0
		}

		for _, value := range array {
			digit := int((uint64(value) / exp) % 10)
			count[digit]++
		}

		sum := 0
		for i := 0; i < 10; i++ {
			old := count[i]
			count[i] = sum
			sum += old
		}

		for _, value := range array {
			digit := int((uint64(value) / exp) % 10)
			pos := count[digit]
			out[pos] = value
			count[digit]++
		}

		copy(array, out)
	}
	return array
}

func absToUint64[I Integers](v I) uint64 {
	x := int64(v)
	if x >= 0 {
		return uint64(x)
	}

	return ^uint64(x) + 1
}
