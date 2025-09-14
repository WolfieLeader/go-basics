package utils

func Reverse(s string) string {
	r := []rune(s)
	for start, end := 0, len(r)-1; start < end; start, end = start+1, end-1 {
		r[start], r[end] = r[end], r[start]
	}
	return string(r)
}

func privateFunction() string {
	return "This is a private function"
}
