package matrix

import "cmp"

type matrixArray[T cmp.Ordered] interface {
	Get(row int, col int) (T, bool)
	GetRow(row int) []T
	GetCol(col int) []T
	Set(row int, col int, value T) bool
	Replace(values ...[]T) bool
	Fill(value T)
	Clear()

	Rows() int
	Cols() int
	Dimensions() (int, int)
	IsEmpty() bool
	Search(value T) (int, int)
	Contains(value T) bool
	Traverse(fn func(row int, col int, value T))
	Swap(row1, col1, row2, col2 int) bool
	SwapRow(row1, row2 int) bool
	SwapCol(col1, col2 int) bool

	Equal(other *MatrixArray[T]) bool
	Copy() *MatrixArray[T]
	String() string
}

var _ matrixArray[int] = (*MatrixArray[int])(nil)
