package dynamic

import "cmp"

type dynamicArray[T cmp.Ordered] interface {
	Get(index int) (T, bool)
	Set(index int, value T) bool
	Replace(values ...T)
	Append(values ...T)
	Prepend(values ...T)
	Insert(index int, value T) bool
	Delete(index int) (T, bool)
	Fill(value T)
	Clear()

	Length() int
	Capacity() int
	IsEmpty() bool
	IsSorted() bool
	Search(value T) int
	BinarySearch(value T) int
	Contains(value T) bool
	Traverse(fn func(index int, value T))
	Swap(index1, index2 int) bool

	Between(start, end int) *DynamicArray[T]
	Copy() *DynamicArray[T]
	Reverse() *DynamicArray[T]
	ToSlice() []T
	Equal(other *DynamicArray[T]) bool
	String() string
}

var _ dynamicArray[int] = (*DynamicArray[int])(nil)
