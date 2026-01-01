package hashmap

type hashMap[K comparable, V comparable] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Delete(key K) (V, bool)
	Contains(key K) bool
	ContainValue(value V) bool
	Size() int
	IsEmpty() bool
	Clear()
	ToMap() map[K]V
	Keys() []K
	Values() []V
	Equal(other *HashMap[K, V]) bool
	Traverse(func(key K, value V) bool)
	Copy() *HashMap[K, V]
	String() string
}

var _ hashMap[string, int] = (*HashMap[string, int])(nil)
