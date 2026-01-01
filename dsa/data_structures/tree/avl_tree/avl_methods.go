package avl

import (
	"fmt"
	"math"
	"strings"
)

func (t *AVLTree[T]) Size() int     { return t.size }
func (t *AVLTree[T]) IsEmpty() bool { return t.size == 0 }
func (t *AVLTree[T]) Clear()        { *t = AVLTree[T]{} }
func (t *AVLTree[T]) Root() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	return t.root.Value, true
}

func (t *AVLTree[T]) Contains(value T) bool {
	curr := t.root
	for curr != nil {
		switch {
		case value < curr.Value:
			curr = curr.left
		case value > curr.Value:
			curr = curr.right
		default: // Equal
			return true
		}
	}
	return false
}

func (n *Node[T]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node[T]) updateHeight() {
	if n == nil {
		return
	}
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *Node[T]) balanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.getHeight() - n.right.getHeight()
}

//     X(2)                               3
//     / \                               / \
//    1   3       Left Rotation        X(2) 5
//       / \      ------------>        / \   \
//      4   5                         1   4   6
//           \
//            6

func (n *Node[T]) leftRotate() *Node[T] {
	if n == nil || n.right == nil {
		return n
	}

	pivot := n.right
	t2 := pivot.left

	pivot.left = n
	n.right = t2

	n.updateHeight()
	pivot.updateHeight()
	return pivot
}

//	       5                         4
//	      / \                       / \
//	     4   6   Right Rotation    3   5
//	    / \      ------------>    /   / \
//	   3   2                    X(1) 2   6
//	  /
//  X(1)

func (n *Node[T]) rightRotate() *Node[T] {
	if n == nil || n.left == nil {
		return n
	}

	pivot := n.left
	t2 := pivot.right

	pivot.right = n
	n.left = t2

	n.updateHeight()
	pivot.updateHeight()
	return pivot
}

func (n *Node[T]) rebalance() *Node[T] {
	if n == nil {
		return nil
	}

	bf := n.balanceFactor()
	if bf > 1 {
		if n.left.balanceFactor() < 0 {
			n.left = n.left.leftRotate()
		}
		return n.rightRotate()
	}

	if bf < -1 {
		if n.right.balanceFactor() > 0 {
			n.right = n.right.rightRotate()
		}
		return n.leftRotate()
	}

	n.updateHeight()
	return n
}

func (t *AVLTree[T]) Insert(values ...T) int {
	inserts := 0
	for _, v := range values {
		var inserted bool
		if t.root, inserted = t.root.insert(v); inserted {
			t.size++
			inserts++
		}
	}
	return inserts
}

func (n *Node[T]) insert(value T) (*Node[T], bool) {
	if n == nil {
		return &Node[T]{Value: value, height: 1}, true
	}

	var ok bool
	switch {
	case value < n.Value:
		n.left, ok = n.left.insert(value)
	case value > n.Value:
		n.right, ok = n.right.insert(value)
	default: // Equal
		return n, false
	}
	return n.rebalance(), ok
}

func (t *AVLTree[T]) Delete(values ...T) int {
	deletes := 0
	for _, v := range values {
		var deleted bool
		if t.root, deleted = t.root.delete(v); deleted {
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
		if n.left, ok = n.left.delete(value); !ok {
			return n, false
		}

	case value > n.Value:
		if n.right, ok = n.right.delete(value); !ok {
			return n, false
		}

	default: // Equal
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
		n.right, _ = n.right.delete(succ.Value)
	}
	return n.rebalance(), true
}

func (t *AVLTree[T]) BalanceFactor() int {
	return t.root.balanceFactor()
}

func (t *AVLTree[T]) Height() int {
	return t.root.getHeight()
}

func (t *AVLTree[T]) Min() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	curr := t.root
	for curr.left != nil {
		curr = curr.left
	}
	return curr.Value, true
}

func (t *AVLTree[T]) Max() (T, bool) {
	var zero T
	if t.root == nil {
		return zero, false
	}
	curr := t.root
	for curr.right != nil {
		curr = curr.right
	}
	return curr.Value, true
}

func (t *AVLTree[T]) Copy() *AVLTree[T] {
	if t.root == nil {
		return New[T]()
	}
	return &AVLTree[T]{root: t.root.copy(), size: t.size}
}

func (n *Node[T]) copy() *Node[T] {
	if n == nil {
		return nil
	}
	return &Node[T]{Value: n.Value, left: n.left.copy(), right: n.right.copy(), height: n.height}
}

func (t *AVLTree[T]) Equal(other *AVLTree[T]) bool {
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
	if n.Value != other.Value || n.height != other.height {
		return false
	}
	return n.left.equal(other.left) && n.right.equal(other.right)
}

func (t *AVLTree[T]) ToSlice() []T {
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

func (t *AVLTree[T]) Successor(value T) (T, bool) { return t.root.successor(value, nil) }
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

func (t *AVLTree[T]) Predecessor(value T) (T, bool) { return t.root.predecessor(value, nil) }
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

func (t *AVLTree[T]) TraverseInOrder(fn func(value T)) { t.root.inOrder(fn) }
func (n *Node[T]) inOrder(fn func(value T)) {
	if n == nil {
		return
	}
	n.left.inOrder(fn)
	fn(n.Value)
	n.right.inOrder(fn)
}

func (t *AVLTree[T]) TraversePreOrder(fn func(value T)) { t.root.preOrder(fn) }
func (n *Node[T]) preOrder(fn func(value T)) {
	if n == nil {
		return
	}
	fn(n.Value)
	n.left.preOrder(fn)
	n.right.preOrder(fn)
}

func (t *AVLTree[T]) TraversePostOrder(fn func(value T)) { t.root.postOrder(fn) }
func (n *Node[T]) postOrder(fn func(value T)) {
	if n == nil {
		return
	}
	n.left.postOrder(fn)
	n.right.postOrder(fn)
	fn(n.Value)
}

func (t *AVLTree[T]) TraverseBreadthFirst(fn func(value T)) {
	if t.root == nil {
		return
	}
	queue := newQueue(t.root)
	for !queue.IsEmpty() {
		n := queue.Dequeue()
		fn(n.Value)
		if n.left != nil {
			queue.Enqueue(n.left)
		}
		if n.right != nil {
			queue.Enqueue(n.right)
		}
	}
}

func (t *AVLTree[T]) IsBalanced() bool { _, ok := t.root.balanced(); return ok }
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

func (t *AVLTree[T]) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "AVL{size=%d}\n", t.size)
	if t.root == nil {
		return sb.String()
	}
	t.root.draw(&sb, "", "", true)
	return sb.String()
}

func (n *Node[T]) draw(sb *strings.Builder, prefix string, label string, isTail bool) {
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
		fmt.Fprintf(sb, "%v(%s%d)\n", n.Value, label, n.height)
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
		child.node.draw(sb, nextPrefix, child.label, i == len(children)-1)
	}
}
