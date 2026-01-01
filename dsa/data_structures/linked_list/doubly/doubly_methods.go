package doubly

import (
	"fmt"
	"strings"
)

func (l DoublyLinkedList[T]) Size() int     { return l.size }
func (l DoublyLinkedList[T]) IsEmpty() bool { return l.size == 0 }
func (l *DoublyLinkedList[T]) Clear()       { *l = DoublyLinkedList[T]{} }

func (l DoublyLinkedList[T]) Copy() *DoublyLinkedList[T] {
	if l.size == 0 {
		return New[T]()
	}

	out := New[T]()
	l.Traverse(func(index int, value T) { out.AddLast(value) })
	return out
}

func (l *DoublyLinkedList[T]) AddFirst(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: l.head}

		if l.head == nil {
			l.tail = n
		} else {
			l.head.prev = n
		}

		l.head = n
		l.size++
	}
}

func (l *DoublyLinkedList[T]) AddLast(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, prev: l.tail}

		if l.tail == nil {
			l.head, l.tail = n, n
			l.size++
			continue
		}

		l.tail.next = n
		l.tail = n
		l.size++
	}
}

func (l DoublyLinkedList[T]) GetFirst() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}
	return l.head.Value, true
}

func (l DoublyLinkedList[T]) GetFirstNode() *Node[T] {
	return l.head
}

func (l DoublyLinkedList[T]) GetLast() (T, bool) {
	var zero T
	if l.tail == nil {
		return zero, false
	}
	return l.tail.Value, true
}

func (l DoublyLinkedList[T]) GetLastNode() *Node[T] {
	return l.tail
}

func (l *DoublyLinkedList[T]) DeleteFirst() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}

	value := l.head.Value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	l.size--
	return value, true
}

func (l *DoublyLinkedList[T]) DeleteLast() (T, bool) {
	var zero T
	if l.tail == nil {
		return zero, false
	}

	value := l.tail.Value
	l.tail = l.tail.prev

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	l.size--
	return value, true
}

func (l DoublyLinkedList[T]) shortestPath(index int) *Node[T] {
	var curr *Node[T]
	if index <= (l.size-1)/2 {
		curr = l.head
		for curr != nil && index > 0 {
			curr = curr.next
			index--
		}
		return curr
	}

	curr = l.tail
	for curr != nil && index < l.size-1 {
		curr = curr.prev
		index++
	}
	return curr
}

func (l *DoublyLinkedList[T]) SetAt(index int, value T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	curr := l.shortestPath(index)
	if curr == nil {
		return false
	}

	curr.Value = value
	return true
}

func (l *DoublyLinkedList[T]) SetAtNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	curr.Value = value
	return true
}

func (l *DoublyLinkedList[T]) InsertAt(index int, value T) bool {
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		l.AddFirst(value)
		return true
	}
	if index == l.size {
		l.AddLast(value)
		return true
	}

	curr := l.shortestPath(index)
	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: curr.prev, next: curr}
	curr.prev.next = n
	curr.prev = n

	l.size++
	return true
}

func (l *DoublyLinkedList[T]) InsertAtNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: curr.prev, next: curr}
	curr.prev.next = n
	curr.prev = n

	l.size++
	return true
}

func (l *DoublyLinkedList[T]) InsertAfter(index int, value T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	if index == l.size-1 {
		l.AddLast(value)
		return true
	}

	curr := l.shortestPath(index)
	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: curr, next: curr.next}
	curr.next = n

	if n.next == nil {
		l.tail = n
	} else {
		n.next.prev = n
	}

	l.size++
	return true
}

func (l *DoublyLinkedList[T]) InsertAfterNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, prev: curr, next: curr.next}
	curr.next = n

	if n.next == nil {
		l.tail = n
	} else {
		n.next.prev = n
	}

	l.size++
	return true
}

func (l *DoublyLinkedList[T]) unlink(curr *Node[T]) (T, bool) {
	var zero T
	if curr == nil {
		return zero, false
	}

	if curr.prev == nil {
		l.head = curr.next
	} else {
		curr.prev.next = curr.next
	}

	if curr.next == nil {
		l.tail = curr.prev
	} else {
		curr.next.prev = curr.prev
	}

	l.size--
	return curr.Value, true
}

func (l *DoublyLinkedList[T]) DeleteAt(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, false
	}

	if index == 0 {
		return l.DeleteFirst()
	}
	if index == l.size-1 {
		return l.DeleteLast()
	}

	curr := l.shortestPath(index)
	return l.unlink(curr)
}

func (l *DoublyLinkedList[T]) DeleteAfter(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, false
	}

	curr := l.shortestPath(index)
	if curr == nil {
		return zero, false
	}

	return l.unlink(curr.next)
}

func (l *DoublyLinkedList[T]) DeleteAtNode(node *Node[T]) (T, bool) {
	var zero T
	if node == nil {
		return zero, false
	}

	if l.head == node {
		return l.DeleteFirst()
	}
	if l.tail == node {
		return l.DeleteLast()
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	return l.unlink(node)
}

func (l *DoublyLinkedList[T]) DeleteAfterNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil {
		return zero, false
	}

	return l.unlink(curr.next)
}

func (l *DoublyLinkedList[T]) DeleteValue(value T) bool {
	if l.head == nil {
		return false
	}

	if l.head.Value == value {
		l.DeleteFirst()
		return true
	}
	if l.tail.Value == value {
		l.DeleteLast()
		return true
	}

	curr := l.head
	for curr != nil {
		if curr.Value == value {
			l.unlink(curr)
			return true
		}

		curr = curr.next
	}
	return false
}

func (l DoublyLinkedList[T]) Get(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, false
	}

	curr := l.shortestPath(index)
	if curr == nil {
		return zero, false
	}

	return curr.Value, true
}

func (l DoublyLinkedList[T]) GetNode(index int) *Node[T] {
	if index < 0 || index >= l.size {
		return nil
	}

	return l.shortestPath(index)

}

func (l DoublyLinkedList[T]) ToSlice() []T {
	out := make([]T, 0, l.size)
	l.Traverse(func(index int, value T) { out = append(out, value) })
	return out
}

func (l DoublyLinkedList[T]) Search(value T) int {
	i := 0
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.Value == value {
			return i
		}
		i++
	}
	return -1
}

func (l DoublyLinkedList[T]) Contains(value T) bool {
	return l.Search(value) != -1
}

func (l *DoublyLinkedList[T]) Equal(other *DoublyLinkedList[T]) bool {
	if l == other {
		return true
	}
	if l == nil || other == nil {
		return false
	}
	if l.size != other.size {
		return false
	}
	curr1, curr2 := l.head, other.head
	for curr1 != nil && curr2 != nil {
		if curr1.Value != curr2.Value {
			return false
		}
		curr1, curr2 = curr1.next, curr2.next
	}
	return curr1 == nil && curr2 == nil
}

func (l DoublyLinkedList[T]) Traverse(fn func(index int, value T)) {
	i := 0
	for curr := l.head; curr != nil; curr = curr.next {
		fn(i, curr.Value)
		i++
	}
}

func (l DoublyLinkedList[T]) Reverse() *DoublyLinkedList[T] {
	out := l.Copy()
	curr := out.head
	for curr != nil {
		curr.prev, curr.next = curr.next, curr.prev
		curr = curr.prev
	}
	out.head, out.tail = out.tail, out.head
	return out
}

func (l DoublyLinkedList[T]) IsSorted() bool {
	if l.size <= 1 {
		return true
	}

	curr := l.head
	for curr != nil && curr.next != nil {
		if curr.Value > curr.next.Value {
			return false
		}
		curr = curr.next
	}
	return true
}

func (l *DoublyLinkedList[T]) Swap(index1, index2 int) bool {
	if index1 < 0 || index2 < 0 || index1 >= l.size || index2 >= l.size {
		return false
	}

	if index1 == index2 {
		return true
	}

	if index1 > index2 {
		index1, index2 = index2, index1
	}

	i := 0
	var n1, n2 *Node[T]
	for curr := l.head; curr != nil; curr = curr.next {
		if i == index1 {
			n1 = curr
		}
		if i == index2 {
			n2 = curr
			break
		}

		i++
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if n1.prev == nil {
		l.head = n2
		n2.prev = nil
	} else {
		n1.prev.next = n2
		n2.prev = n1.prev
	}

	if n1.next == n2 {
		n1.next = n2.next

		if n2.next == nil {
			l.tail = n1
		} else {
			n2.next.prev = n1
		}

		n2.next = n1
		n1.prev = n2

		return true
	}

	n2.prev.next = n1
	n1.prev = n2.prev

	realNext1, realNext2 := n1.next, n2.next

	if realNext1 == nil {
		l.tail = n2
	} else {
		realNext1.prev = n2
	}

	if realNext2 == nil {
		l.tail = n1
	} else {
		realNext2.prev = n1
	}

	n1.next, n2.next = realNext2, realNext1

	return true
}

func (l DoublyLinkedList[T]) String() string {
	if l.head == nil {
		return "[nil]"
	}

	var sb strings.Builder
	sb.WriteString("[nil <- (head) ")

	for curr := l.head; curr != nil; curr = curr.next {
		fmt.Fprintf(&sb, "%v", curr.Value)
		if curr.next != nil {
			sb.WriteString(" <-> ")
		} else {
			sb.WriteString(" (tail) -> nil]")
		}
	}
	return sb.String()
}
