package bst

func (t *BinarySearchTree[T]) Size() int     { return t.size }
func (t *BinarySearchTree[T]) IsEmpty() bool { return t.size == 0 }
func (t *BinarySearchTree[T]) Clear()        { *t = BinarySearchTree[T]{} }
func (t *BinarySearchTree[T]) Root() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	return t.root.Value, true
}
