package bst

import (
	"fmt"
	"math"
	"strings"
)

func (t *BinarySearchTree[T]) ContainsR(value T) bool { return t.root.Contains(value) }
func (n *Node[T]) Contains(value T) bool {
	if n == nil {
		return false
	}

	switch {
	case value < n.Value:
		return n.left.Contains(value)
	case value > n.Value:
		return n.right.Contains(value)
	default: // Equal
		return true
	}
}

func (t *BinarySearchTree[T]) InsertR(values ...T) {
	for _, v := range values {
		var ok bool
		if t.root, ok = t.root.insert(v); ok {
			t.size++
		}
	}
}

func (n *Node[T]) insert(value T) (*Node[T], bool) {
	if n == nil {
		return &Node[T]{Value: value}, true
	}

	var ok bool
	switch {
	case value < n.Value:
		n.left, ok = n.left.insert(value)
	case value > n.Value:
		n.right, ok = n.right.insert(value)
	default: // duplicate
		return n, false
	}
	return n, ok
}

func (t *BinarySearchTree[T]) DeleteR(values ...T) int {
	deletes := 0
	for _, v := range values {
		var ok bool
		if t.root, ok = t.root.delete(v); ok {
			t.size--
			deletes++
		}
	}
	return deletes
}

func (n *Node[T]) delete(value T) (*Node[T], bool) {
	if n == nil {
		return nil, false
	}

	var ok bool
	switch {
	case value < n.Value:
		n.left, ok = n.left.delete(value)
	case value > n.Value:
		n.right, ok = n.right.delete(value)
	default: //Equal
		if n.left == nil {
			return n.right, true
		}
		if n.right == nil {
			return n.left, true
		}

		succ := n.right
		for succ.left != nil {
			succ = succ.left
		}
		n.Value = succ.Value
		n.right, ok = n.right.delete(succ.Value)
	}
	return n, ok
}

func (t *BinarySearchTree[T]) MinR() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	return t.root.min(), true
}

func (n *Node[T]) min() T {
	if n.left == nil {
		return n.Value
	}
	return n.left.min()
}

func (t *BinarySearchTree[T]) MaxR() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	return t.root.max(), true
}

func (n *Node[T]) max() T {
	if n.right == nil {
		return n.Value
	}
	return n.right.max()
}

func (t *BinarySearchTree[T]) TraverseInOrderR(fn func(value T)) { t.root.inOrder(fn) }
func (n *Node[T]) inOrder(fn func(value T)) {
	if n == nil {
		return
	}
	n.left.inOrder(fn)
	fn(n.Value)
	n.right.inOrder(fn)
}

func (t *BinarySearchTree[T]) TraversePreOrderR(fn func(value T)) { t.root.preOrder(fn) }
func (n *Node[T]) preOrder(fn func(value T)) {
	if n == nil {
		return
	}
	fn(n.Value)
	n.left.preOrder(fn)
	n.right.preOrder(fn)
}

func (t *BinarySearchTree[T]) TraversePostOrderR(fn func(value T)) { t.root.postOrder(fn) }
func (n *Node[T]) postOrder(fn func(value T)) {
	if n == nil {
		return
	}
	n.left.postOrder(fn)
	n.right.postOrder(fn)
	fn(n.Value)
}

func (t *BinarySearchTree[T]) HeightR() int { return t.root.height() }
func (n *Node[T]) height() int {
	if n == nil {
		return 0
	}
	return max(n.left.height(), n.right.height()) + 1
}

func (t *BinarySearchTree[T]) SuccessorR(value T) (T, bool) { return t.root.successor(value, nil) }
func (n *Node[T]) successor(value T, succ *Node[T]) (T, bool) {
	var zero T
	if n == nil {
		if succ != nil {
			return succ.Value, true
		}
		return zero, false
	}

	switch {
	case value < n.Value:
		return n.left.successor(value, n)
	case value > n.Value:
		return n.right.successor(value, succ)
	default: // Equal
		if r := n.right; r != nil {
			for r.left != nil {
				r = r.left
			}
			return r.Value, true
		}

		if succ != nil {
			return succ.Value, true
		}

		return zero, false
	}
}

func (t *BinarySearchTree[T]) PredecessorR(value T) (T, bool) { return t.root.predecessor(value, nil) }
func (n *Node[T]) predecessor(value T, pred *Node[T]) (T, bool) {
	var zero T
	if n == nil {
		if pred != nil {
			return pred.Value, true
		}
		return zero, false
	}

	switch {
	case value < n.Value:
		return n.left.predecessor(value, pred)
	case value > n.Value:
		return n.right.predecessor(value, n)
	default: // Equal
		if l := n.left; l != nil {
			for l.right != nil {
				l = l.right
			}
			return l.Value, true
		}

		if pred != nil {
			return pred.Value, true
		}

		return zero, false
	}
}

func (t *BinarySearchTree[T]) IsBalanced() bool { _, ok := t.root.balanced(); return ok }
func (n *Node[T]) balanced() (int, bool) {
	if n == nil {
		return 0, true
	}
	leftHeight, leftBalanced := n.left.balanced()
	if !leftBalanced {
		return 0, false
	}
	rightHeight, rightBalanced := n.right.balanced()
	if !rightBalanced {
		return 0, false
	}

	height := max(leftHeight, rightHeight) + 1
	isBalanced := math.Abs(float64(leftHeight-rightHeight)) <= 1
	return height, isBalanced
}

func (t *BinarySearchTree[T]) EqualR(other *BinarySearchTree[T]) bool {
	if t == other {
		return true
	}
	if t == nil || other == nil {
		return false
	}
	if t.size != other.size {
		return false
	}
	return t.root.equal(other.root)
}
func (n *Node[T]) equal(other *Node[T]) bool {
	if n == nil || other == nil {
		return n == nil && other == nil
	}
	if n.Value != other.Value {
		return false
	}
	return n.left.equal(other.left) && n.right.equal(other.right)
}

func (t *BinarySearchTree[T]) ToSliceR() []T {
	if t.root == nil {
		return nil
	}
	out := make([]T, 0, t.size)
	t.root.toSlice(&out)
	return out
}
func (n *Node[T]) toSlice(out *[]T) {
	if n == nil {
		return
	}
	n.left.toSlice(out)
	*out = append(*out, n.Value)
	n.right.toSlice(out)
}

func (t *BinarySearchTree[T]) CopyR() *BinarySearchTree[T] {
	if t.root == nil {
		return New[T]()
	}
	return &BinarySearchTree[T]{root: t.root.copy(), size: t.size}
}

func (n *Node[T]) copy() *Node[T] {
	if n == nil {
		return nil
	}
	return &Node[T]{Value: n.Value, left: n.left.copy(), right: n.right.copy()}
}

// Cannot be used in BST only in Binary Tree
// func (t *BST[T]) IsSymmetric() bool {
// 	if t.root == nil {
// 		return true
// 	}
// 	return t.root.left.symmetric(t.root.right)
// }

// func (n *Node[T]) symmetric(other *Node[T]) bool {
// 	if n == nil || other == nil {
// 		return n == nil && other == nil
// 	}
// 	if n.Value != other.Value {
// 		return false
// 	}
// 	return n.left.symmetric(other.right) && n.right.symmetric(other.left)
// }

func (t *BinarySearchTree[T]) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "BST{size=%d}\n", t.size)
	if t.root == nil {
		return sb.String()
	}
	t.root.draw(&sb, "", 0, "", true)
	return sb.String()
}

func (n *Node[T]) draw(sb *strings.Builder, prefix string, level int, label string, isTail bool) {
	if n == nil {
		return
	}

	sb.WriteString(prefix)
	switch {
	case prefix == "":
		sb.WriteString("└── ")
	case isTail:
		sb.WriteString("└── ")
	default:
		sb.WriteString("├── ")
	}

	if prefix == "" {
		fmt.Fprintf(sb, "%v\n", n.Value)
	} else {
		fmt.Fprintf(sb, "%v(%s%d)\n", n.Value, label, level)
	}

	type labeled struct {
		node  *Node[T]
		label string
	}

	children := make([]labeled, 0, 2)
	if n.right != nil {
		children = append(children, labeled{n.right, "R"})
	}
	if n.left != nil {
		children = append(children, labeled{n.left, "L"})
	}

	for i, child := range children {
		var nextPrefix string
		switch {
		case prefix == "":
			nextPrefix = "    "
		case isTail:
			nextPrefix = prefix + "    "
		default:
			nextPrefix = prefix + "│   "
		}
		child.node.draw(sb, nextPrefix, level+1, child.label, i == len(children)-1)
	}
}
