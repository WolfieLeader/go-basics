package hashmap

type HashMap[K comparable, V comparable] struct {
	data map[K]V
}

func New[K comparable, V comparable]() *HashMap[K, V] {
	return &HashMap[K, V]{data: make(map[K]V)}
}
