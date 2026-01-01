package hashset

type hashSet[T comparable] interface {
	Add(values ...T)
	Delete(values ...T) int
	Contains(value T) bool
	Size() int
	IsEmpty() bool
	Clear()
	ToSlice() []T
	Equal(other *HashSet[T]) bool
	Traverse(func(value T) bool)
	Copy() *HashSet[T]
	String() string
}

var _ hashSet[string] = (*HashSet[string])(nil)
