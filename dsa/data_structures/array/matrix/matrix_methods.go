package matrix

import (
	"cmp"
	"fmt"
	"slices"
)

func (a *MatrixArray[T]) Rows() int              { return a.rows }
func (a *MatrixArray[T]) Cols() int              { return a.cols }
func (a *MatrixArray[T]) Dimensions() (int, int) { return a.rows, a.cols }
func (a *MatrixArray[T]) IsEmpty() bool          { return a.rows == 0 || a.cols == 0 }
func (a *MatrixArray[T]) Clear() {
	var zero T
	a.Fill(zero)
}

func (a *MatrixArray[T]) Get(row int, col int) (T, bool) {
	if row < 0 || row >= a.rows || col < 0 || col >= a.cols {
		var zero T
		return zero, false
	}
	return a.data[row][col], true
}

func (a *MatrixArray[T]) GetRow(row int) []T {
	if row < 0 || row >= a.rows {
		return nil
	}

	out := make([]T, a.cols)
	copy(out, a.data[row])
	return out
}

func (a *MatrixArray[T]) GetCol(col int) []T {
	if col < 0 || col >= a.cols {
		return nil
	}

	out := make([]T, a.rows)
	for i := 0; i < a.rows; i++ {
		out[i] = a.data[i][col]
	}
	return out
}

func (a *MatrixArray[T]) Set(row int, col int, value T) bool {
	if row < 0 || row >= a.rows || col < 0 || col >= a.cols {
		return false
	}

	a.data[row][col] = value
	return true
}

func (a *MatrixArray[T]) Replace(values ...[]T) bool {
	if len(values) > a.rows {
		return false
	}

	for _, row := range values {
		if len(row) > a.cols {
			return false
		}
	}

	var zero T
	a.Fill(zero)

	for i := 0; i < len(values); i++ {
		copy(a.data[i], values[i])
	}
	return true
}

func (a *MatrixArray[T]) Fill(value T) {
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			a.data[i][j] = value
		}
	}
}

func (a *MatrixArray[T]) Search(value T) (int, int) {
	for row := 0; row < a.rows; row++ {
		for col := 0; col < a.cols; col++ {
			if cmp.Compare(a.data[row][col], value) == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

func (a *MatrixArray[T]) Contains(value T) bool {
	row, col := a.Search(value)
	return row != -1 && col != -1
}

func (a *MatrixArray[T]) Traverse(fn func(row int, col int, value T)) {
	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			fn(i, j, a.data[i][j])
		}
	}
}

func (a *MatrixArray[T]) Swap(row1, col1, row2, col2 int) bool {
	if row1 < 0 || row1 >= a.rows || col1 < 0 || col1 >= a.cols ||
		row2 < 0 || row2 >= a.rows || col2 < 0 || col2 >= a.cols {
		return false
	}

	a.data[row1][col1], a.data[row2][col2] = a.data[row2][col2], a.data[row1][col1]
	return true
}

func (a *MatrixArray[T]) SwapRow(row1, row2 int) bool {
	if row1 < 0 || row1 >= a.rows || row2 < 0 || row2 >= a.rows {
		return false
	}

	a.data[row1], a.data[row2] = a.data[row2], a.data[row1]
	return true
}

func (a *MatrixArray[T]) SwapCol(col1, col2 int) bool {
	if col1 < 0 || col1 >= a.cols || col2 < 0 || col2 >= a.cols {
		return false
	}

	for i := 0; i < a.rows; i++ {
		a.data[i][col1], a.data[i][col2] = a.data[i][col2], a.data[i][col1]
	}
	return true
}

func (a *MatrixArray[T]) Equal(other *MatrixArray[T]) bool {
	if a == other {
		return true
	}
	if a == nil || other == nil {
		return false
	}
	return slices.EqualFunc(a.data, other.data, func(x, y []T) bool { return slices.Equal(x, y) })
}

func (a *MatrixArray[T]) Copy() *MatrixArray[T] {
	out := make([][]T, a.rows)
	for i := 0; i < a.rows; i++ {
		out[i] = make([]T, a.cols)
		copy(out[i], a.data[i])
	}
	return &MatrixArray[T]{data: out, rows: a.rows, cols: a.cols}
}

func (a *MatrixArray[T]) String() string {
	return fmt.Sprintf("%v", a.data)
}
