package bst

func (t *BinarySearchTree[T]) ContainsI(value T) bool {
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

func (t *BinarySearchTree[T]) InsertI(values ...T) {
	for _, v := range values {
		t.insertI(v)
	}
}

func (t *BinarySearchTree[T]) insertI(value T) bool {
	if t.root == nil {
		t.root = &Node[T]{Value: value}
		t.size = 1
		return true
	}

	curr := t.root
	for {
		switch {
		case value < curr.Value:
			if curr.left == nil {
				curr.left = &Node[T]{Value: value}
				t.size++
				return true
			}
			curr = curr.left

		case value > curr.Value:
			if curr.right == nil {
				curr.right = &Node[T]{Value: value}
				t.size++
				return true
			}
			curr = curr.right

		default: // Equal
			return false
		}
	}
}

func (t *BinarySearchTree[T]) DeleteI(values ...T) int {
	deletes := 0
	for _, v := range values {
		if t.deleteI(v) {
			deletes++
		}
	}
	return deletes
}

func (t *BinarySearchTree[T]) deleteI(value T) bool {
	curr, ptr := t.root, &t.root
outer:
	for curr != nil {
		switch {
		case value < curr.Value:
			ptr = &curr.left
			curr = curr.left

		case value > curr.Value:
			ptr = &curr.right
			curr = curr.right

		default: // Equal
			break outer
		}
	}

	if curr == nil {
		return false
	}

	if curr.left != nil && curr.right != nil {
		succ, succPtr := curr.right, &curr.right
		for succ.left != nil {
			succPtr = &succ.left
			succ = succ.left
		}

		curr.Value = succ.Value
		curr, ptr = succ, succPtr
	}

	if curr.left != nil {
		*ptr = curr.left
	} else {
		*ptr = curr.right
	}

	t.size--
	return true
}

func (t *BinarySearchTree[T]) MinI() (T, bool) {
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

func (t *BinarySearchTree[T]) MaxI() (T, bool) {
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

func (t *BinarySearchTree[T]) TraverseInOrderI(fn func(value T)) {
	if t.root == nil {
		return
	}
	stack := newStack[*Node[T]]()
	curr := t.root
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.left
		}

		curr = stack.Pop()
		fn(curr.Value)
		curr = curr.right
	}
}

func (t *BinarySearchTree[T]) TraversePreOrderI(fn func(value T)) {
	if t.root == nil {
		return
	}
	stack := newStack(t.root)
	for !stack.IsEmpty() {
		n := stack.Pop()
		fn(n.Value)
		if n.right != nil {
			stack.Push(n.right)
		}
		if n.left != nil {
			stack.Push(n.left)
		}
	}
}

func (t *BinarySearchTree[T]) TraversePostOrderI(fn func(value T)) {
	if t.root == nil {
		return
	}
	var visited *Node[T]
	stack := newStack[*Node[T]]()
	curr := t.root
	for curr != nil || !stack.IsEmpty() {
		if curr != nil {
			stack.Push(curr)
			curr = curr.left
			continue
		}

		n := stack.Peek()
		if n.right != nil && n.right != visited {
			curr = n.right
			continue
		}

		fn(n.Value)
		visited = n
		stack.Pop()
	}
}

func (t *BinarySearchTree[T]) TraverseBreadthFirst(fn func(value T)) {
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

func (t *BinarySearchTree[T]) HeightI() int {
	if t.root == nil {
		return 0
	}
	queue := newQueue(t.root)
	height := 0
	for !queue.IsEmpty() {
		levelSize := queue.Size()
		for range levelSize {
			n := queue.Dequeue()
			if n.left != nil {
				queue.Enqueue(n.left)
			}
			if n.right != nil {
				queue.Enqueue(n.right)
			}
		}
		height++
	}
	return height
}

func (t *BinarySearchTree[T]) SuccessorI(value T) (T, bool) {
	var zero T
	var succ *Node[T]

	curr := t.root
	for curr != nil {
		switch {
		case value < curr.Value:
			succ = curr
			curr = curr.left

		case value > curr.Value:
			curr = curr.right

		default: // Equal
			if r := curr.right; r != nil {
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
	if succ != nil {
		return succ.Value, true
	}
	return zero, false
}

func (t *BinarySearchTree[T]) PredecessorI(value T) (T, bool) {
	var zero T
	var pred *Node[T]

	curr := t.root
	for curr != nil {
		switch {
		case value < curr.Value:
			curr = curr.left

		case value > curr.Value:
			pred = curr
			curr = curr.right

		default: // Equal
			if l := curr.left; l != nil {
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
	if pred != nil {
		return pred.Value, true
	}
	return zero, false
}

func (t *BinarySearchTree[T]) EqualI(other *BinarySearchTree[T]) bool {
	if t == other {
		return true
	}
	if t == nil || other == nil {
		return false
	}
	if t.size != other.size {
		return false
	}

	type pair struct{ t1, t2 *Node[T] }
	stack := newStack(pair{t1: t.root, t2: other.root})

	for !stack.IsEmpty() {
		n := stack.Pop()
		if n.t1.Value != n.t2.Value {
			return false
		}
		l1, l2 := n.t1.left, n.t2.left
		if (l1 == nil) != (l2 == nil) {
			return false
		}
		if l1 != nil {
			stack.Push(pair{t1: l1, t2: l2})
		}
		r1, r2 := n.t1.right, n.t2.right
		if (r1 == nil) != (r2 == nil) {
			return false
		}
		if r1 != nil {
			stack.Push(pair{t1: r1, t2: r2})
		}
	}
	return true
}

func (t *BinarySearchTree[T]) ToSliceI() []T {
	if t.root == nil {
		return nil
	}

	out := make([]T, 0, t.size)
	stack := newStack[*Node[T]]()
	curr := t.root
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.left
		}

		curr = stack.Pop()
		out = append(out, curr.Value)
		curr = curr.right
	}
	return out
}

func (t *BinarySearchTree[T]) CopyI() *BinarySearchTree[T] {
	if t.root == nil {
		return New[T]()
	}
	newTree := &BinarySearchTree[T]{root: &Node[T]{Value: t.root.Value}, size: t.size}

	type pair struct{ src, dst *Node[T] }
	stack := newStack(pair{src: t.root, dst: newTree.root})
	for !stack.IsEmpty() {
		n := stack.Pop()
		if right := n.src.right; right != nil {
			n.dst.right = &Node[T]{Value: right.Value}
			stack.Push(pair{src: right, dst: n.dst.right})
		}
		if left := n.src.left; left != nil {
			n.dst.left = &Node[T]{Value: left.Value}
			stack.Push(pair{src: left, dst: n.dst.left})
		}
	}
	return newTree
}
