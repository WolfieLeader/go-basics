package hashset

type HashSet[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable]() *HashSet[T] {
	return &HashSet[T]{data: make(map[T]struct{})}
}
