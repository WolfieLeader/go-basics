package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	avl "github.com/WolfieLeader/go-basics/dsa/data_structures/tree/avl_tree"
	bst "github.com/WolfieLeader/go-basics/dsa/data_structures/tree/binary_search_tree"
	heap "github.com/WolfieLeader/go-basics/dsa/data_structures/tree/heap"
)

func main() {
	fmt.Println("Binary Search Tree Example")
	bstExample()
	fmt.Println()

	fmt.Println("AVL Tree Example")
	avlExample()
	fmt.Println()

	fmt.Println("Heap Example")
	heapExample()
	fmt.Println()
}

func bstExample() {
	t := bst.New[int]()
	t.InsertI(10, 5, 15, 2, 5)
	t.InsertR(7, 13, 17, 10)
	fmt.Printf("- Tree:\n%v", t)

	root, _ := t.Root()
	minI, _ := t.MinI()
	maxI, _ := t.MaxI()
	minR, _ := t.MinR()
	maxR, _ := t.MaxR()
	fmt.Printf("- Size: %d, Root: %v, Height: %d(ok=%t)\n", t.Size(), root, t.HeightI(), t.HeightR() == t.HeightI())
	fmt.Printf("- Min: %v(ok=%t), Max: %v(ok=%t), Is Balanced: %t\n", minI, minI == minR, maxI, maxI == maxR, t.IsBalanced())
	fmt.Printf("- Contains 17: %t(ok=%t)\n", t.ContainsI(17), t.ContainsI(17) == t.ContainsR(17))
	fmt.Printf("- Contains 0: %t(ok=%t)\n", t.ContainsI(0), t.ContainsI(0) == t.ContainsR(0))

	dfsInOrderI, dfsInOrderR := traverse(t.TraverseInOrderI), traverse(t.TraverseInOrderR)
	dfsPreOrderI, dfsPreOrderR := traverse(t.TraversePreOrderI), traverse(t.TraversePreOrderR)
	dfsPostOrderI, dfsPostOrderR := traverse(t.TraversePostOrderI), traverse(t.TraversePostOrderR)
	bfs := traverse(t.TraverseBreadthFirst)
	fmt.Printf("- DFS In Order: %s(ok=%t)\n", dfsInOrderI, dfsInOrderI == dfsInOrderR)
	fmt.Printf("- DFS Pre Order: %s(ok=%t)\n", dfsPreOrderI, dfsPreOrderI == dfsPreOrderR)
	fmt.Printf("- DFS Post Order: %s(ok=%t)\n", dfsPostOrderI, dfsPostOrderI == dfsPostOrderR)
	fmt.Printf("- BFS: %s\n", bfs)

	for _, v := range []int{6, 10, 13, 100, -1} {
		pi, piOk := t.PredecessorI(v)
		pr, prOk := t.PredecessorR(v)
		if piOk && prOk {
			fmt.Printf("- %d Predecessor: %v(ok=%t), ", v, pi, pi == pr)
		} else {
			fmt.Printf("- %d Predecessor: none, ", v)
		}

		si, siOk := t.SuccessorI(v)
		sr, srOk := t.SuccessorR(v)
		if siOk && srOk {
			fmt.Printf("%d Successor: %v(ok=%t)\n", v, si, si == sr)
		} else {
			fmt.Printf("%d Successor: none\n", v)
		}
	}

	fmt.Printf("- Delete 17, 5, 10, 2: %v\n", t.DeleteI(17, 5) == t.DeleteR(10, 2))
	fmt.Printf("- After deletes:\n%v", t)
	fmt.Printf("- Size: %d, Height: %d(ok=%t), Is Balanced: %t\n", t.Size(), t.HeightI(), t.HeightR() == t.HeightI(), t.IsBalanced())
	t.Clear()
	fmt.Printf("- After clear:\n%v", t)
	fmt.Printf("- Size: %d, IsEmpty: %t, Is Balanced: %t\n", t.Size(), t.IsEmpty(), t.IsBalanced())

	t.InsertI(1, 2, 3)
	t.InsertR(4, 5, 6)
	fmt.Printf("- Skewed tree:\n%v", t)
	fmt.Printf("- To Slice: %v(ok=%t)\n", t.ToSliceI(), slices.Equal(t.ToSliceI(), t.ToSliceR()))
	fmt.Printf("- Size: %d, Height: %d(ok=%t), Is Balanced: %t\n", t.Size(), t.HeightI(), t.HeightR() == t.HeightI(), t.IsBalanced())
	c1 := t.CopyI()
	c1.InsertI(0)
	c2 := t.CopyR()
	c2.InsertR(7)
	fmt.Printf("- Are BSTs equal? (skew tree + 0) == (skew tree + 7): %t(ok=%t)\n", c1.EqualI(c2), c1.EqualI(c2) == c1.EqualR(c2))
	c1.DeleteI(0)
	c2.DeleteR(7)
	fmt.Printf("- Are BSTs equal? (skew tree) == (skew tree + 0 - 0) == (skew tree + 7 - 7): %t(ok=%t)\n", t.EqualI(c1) && t.EqualI(c2), t.EqualI(c1) && t.EqualI(c2) == t.EqualR(c1) && t.EqualR(c2))
}

func avlExample() {
	t := avl.New[int]()
	t.Insert(10, 5, 15, 2, 5, 7, 13, 17, 10)
	fmt.Printf("- Tree:\n%v", t)

	root, _ := t.Root()
	min, _ := t.Min()
	max, _ := t.Max()
	fmt.Printf("- Size: %d, Root: %v, Height: %d\n", t.Size(), root, t.Height())
	fmt.Printf("- Min: %v, Max: %v, Is Balanced: %t\n", min, max, t.IsBalanced())
	fmt.Printf("- Contains 17: %t\n", t.Contains(17))
	fmt.Printf("- Contains 0: %t\n", t.Contains(0))

	fmt.Printf("- DFS In Order: %s\n", traverse(t.TraverseInOrder))
	fmt.Printf("- DFS Pre Order: %s\n", traverse(t.TraversePreOrder))
	fmt.Printf("- DFS Post Order: %s\n", traverse(t.TraversePostOrder))
	fmt.Printf("- BFS: %s\n", traverse(t.TraverseBreadthFirst))

	for _, v := range []int{6, 10, 13, 100, -1} {
		if pred, ok := t.Predecessor(v); ok {
			fmt.Printf("- %d Predecessor: %v, ", v, pred)
		} else {
			fmt.Printf("- %d Predecessor: none, ", v)
		}

		if succ, ok := t.Successor(v); ok {
			fmt.Printf("%d Successor: %v\n", v, succ)
		} else {
			fmt.Printf("%d Successor: none\n", v)
		}
	}

	fmt.Printf("- Delete 17, 5, 10, 2: %v\n", t.Delete(17, 5, 10, 2) == 4)
	fmt.Printf("- After deletes:\n%v", t)
	fmt.Printf("- Size: %d, Height: %d, Is Balanced: %t\n", t.Size(), t.Height(), t.IsBalanced())
	t.Clear()
	fmt.Printf("- After clear:\n%v", t)
	fmt.Printf("- Size: %d, IsEmpty: %t, Is Balanced: %t\n", t.Size(), t.IsEmpty(), t.IsBalanced())

	t.Insert(1, 2, 3, 4, 5, 6)
	fmt.Printf("- Test if skewed:\n%v", t)
	fmt.Printf("- To Slice: %v\n", t.ToSlice())
	fmt.Printf("- Size: %d, Height: %d, Is Balanced: %t\n", t.Size(), t.Height(), t.IsBalanced())
	c := t.Copy()
	c.Insert(10)
	fmt.Printf("- Are AVLs equal? (tree) == (tree + 10): %t\n", t.Equal(c))
	c.Delete(10)
	fmt.Printf("- After delete 10:\n%v", c)
	fmt.Printf("- Are AVLs equal? (tree) == (tree + 10 - 10): %t\n", t.Equal(c))
	t.Insert(10)
	t.Delete(10)
	fmt.Printf("- Are AVLs equal? (tree + 10 - 10) == (tree + 10 - 10): %t\n", t.Equal(c))
}

func heapExample() {
	minH := heap.New[int](true)
	minH.Push(5, 3, 8, 10, 4, 50, 6, 8, 3, 4, 2, 0, 1, -3, 7, 9)
	fmt.Printf("- Min-Heap:\n%v", minH)

	peek, _ := minH.Peek()
	fmt.Printf("- Size: %d, IsEmpty: %t, Peek: %v\n", minH.Size(), minH.IsEmpty(), peek)
	fmt.Printf("- BFS: %s\n", traverse(minH.TraverseBreadthFirst))
	fmt.Printf("- To Slice: %v\n", minH.ToSlice())

	popped := make([]int, 0, 5)
	for range 5 {
		if v, ok := minH.Pop(); ok {
			popped = append(popped, v)
		}
	}
	fmt.Printf("- Pop 5: %v\n", popped)
	fmt.Printf("- After pops:\n%v", minH)

	for _, q := range []int{-3, 42, 5} {
		fmt.Printf("- Contains %d: %t\n", q, minH.Contains(q))
	}

	input := []int{7, 1, 9, 0, 5, 2, 8, 4, 6, 3, 11, 12, 21, 19}
	maxH := heap.New[int](false)
	maxH.Heapify(input...)
	fmt.Printf("- Heapify %v:\n%v", input, maxH)

	c := maxH.Copy()
	fmt.Printf("- Equal (maxH, copy): %t\n", maxH.Equal(c))
	_, _ = c.Pop()
	fmt.Printf("- Equal after one pop: %t\n", maxH.Equal(c))

	type person struct {
		name string
		age  int
	}
	byAgeThenName := func(a, b person) bool {
		if a.age != b.age {
			return a.age > b.age
		}
		return a.name < b.name
	}
	people := heap.NewByFunc(byAgeThenName)
	people.Push(person{"Alice", 30}, person{"Bob", 25}, person{"Charlie", 30}, person{"Dave", 20})
	fmt.Printf("- Max-Heap by age, then name:\n%v", people)
	fmt.Printf("- Size: %d, IsEmpty: %t\n", people.Size(), people.IsEmpty())
	fmt.Printf("- Pop order: ")
	for !people.IsEmpty() {
		p, _ := people.Pop()
		fmt.Printf("%s(%d) ", p.name, p.age)
	}
	fmt.Println()
}

func traverse[T cmp.Ordered](method func(fn func(value T))) string {
	var sb strings.Builder
	first := true
	method(func(v T) {
		if !first {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%v", v)
		first = false
	})
	return sb.String()
}
