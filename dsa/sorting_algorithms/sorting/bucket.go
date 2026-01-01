package sorting

func Bucket[T Numbers](arr []T) []T {
	array, length := copyArray(arr)
	if length <= 1 {
		return array
	}

	minVal, maxVal := array.findMinMax()
	if minVal == maxVal {
		return array
	}

	buckets := make([][]T, length)

	minF, maxF := toFloat64(minVal), toFloat64(maxVal)
	rangeF := maxF - minF
	for _, value := range array {
		placement := int(((toFloat64(value) - minF) / rangeF) * float64(length-1))
		index := min(max(placement, 0), length-1)
		buckets[index] = append(buckets[index], value)
	}

	for i := range buckets {
		buckets[i] = Insertion(buckets[i])
	}

	out := make([]T, 0, length)
	for _, bucket := range buckets {
		out = append(out, bucket...)
	}

	return out
}

func toFloat64[T Numbers](x T) float64 {
	switch any(x).(type) {
	case int, int8, int16, int32, int64:
		return float64(int64(x))
	case uint, uint8, uint16, uint32, uint64:
		return float64(uint64(x))
	case float32, float64:
		return float64(x)
	default:
		return 0
	}
}
