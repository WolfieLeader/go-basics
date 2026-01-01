package singly

import (
	"fmt"
	"strings"
)

func (l *SinglyLinkedList[T]) Size() int     { return l.size }
func (l *SinglyLinkedList[T]) IsEmpty() bool { return l.size == 0 }
func (l *SinglyLinkedList[T]) Clear()        { *l = SinglyLinkedList[T]{} }

func (l *SinglyLinkedList[T]) AddFirst(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: l.head}

		if l.head == nil {
			l.tail = n
		}

		l.head = n
		l.size++
	}
}

func (l *SinglyLinkedList[T]) AddLast(values ...T) {
	for _, value := range values {
		n := &Node[T]{Value: value, next: nil}

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

func (l *SinglyLinkedList[T]) GetFirst() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}
	return l.head.Value, true
}

func (l *SinglyLinkedList[T]) GetFirstNode() *Node[T] {
	return l.head
}

func (l *SinglyLinkedList[T]) GetLast() (T, bool) {
	var zero T
	if l.tail == nil {
		return zero, false
	}
	return l.tail.Value, true
}

func (l *SinglyLinkedList[T]) GetLastNode() *Node[T] {
	return l.tail
}

func (l *SinglyLinkedList[T]) DeleteFirst() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}

	value := l.head.Value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	}

	l.size--
	return value, true
}

func (l *SinglyLinkedList[T]) DeleteLast() (T, bool) {
	var zero T
	if l.tail == nil {
		return zero, false
	}

	value := l.tail.Value

	if l.head == l.tail {
		l.Clear()
		return value, true
	}

	prev := l.head
	for prev.next != l.tail {
		prev = prev.next
	}

	prev.next = nil
	l.tail = prev

	l.size--
	return value, true
}

func (l *SinglyLinkedList[T]) SetAt(index int, value T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	curr := l.head
	for curr != nil && index > 0 {
		curr = curr.next
		index--
	}

	if curr == nil {
		return false
	}

	curr.Value = value
	return true
}

func (l *SinglyLinkedList[T]) SetAtNode(node *Node[T], value T) bool {
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

func (l *SinglyLinkedList[T]) InsertAt(index int, value T) bool {
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

	prev := l.head
	for prev != nil && index > 1 {
		prev = prev.next
		index--
	}

	if prev == nil {
		return false
	}

	n := &Node[T]{Value: value, next: prev.next}
	prev.next = n

	if n.next == nil {
		l.tail = n
	}

	l.size++
	return true
}

func (l *SinglyLinkedList[T]) InsertAtNode(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	if l.head == node {
		l.AddFirst(value)
		return true
	}
	if l.tail == node {
		l.AddLast(value)
		return true
	}

	prev := l.head
	for prev != nil && prev.next != node {
		prev = prev.next
	}

	if prev == nil {
		return false
	}

	n := &Node[T]{Value: value, next: prev.next}
	prev.next = n

	if n.next == nil {
		l.tail = n
	}

	l.size++
	return true
}

func (l *SinglyLinkedList[T]) InsertAfter(index int, value T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	if index == l.size-1 {
		l.AddLast(value)
		return true
	}

	curr := l.head
	for curr != nil && index > 0 {
		curr = curr.next
		index--
	}

	if curr == nil {
		return false
	}

	n := &Node[T]{Value: value, next: curr.next}
	curr.next = n

	if n.next == nil {
		l.tail = n
	}

	l.size++
	return true
}

func (l *SinglyLinkedList[T]) InsertAfterNode(node *Node[T], value T) bool {
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

	n := &Node[T]{Value: value, next: curr.next}
	curr.next = n

	if n.next == nil {
		l.tail = n
	}

	l.size++
	return true
}

func (l *SinglyLinkedList[T]) DeleteAt(index int) (T, bool) {
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

	prev := l.head
	for prev != nil && index > 1 {
		prev = prev.next
		index--
	}

	if prev == nil || prev.next == nil {
		return zero, false
	}

	value := prev.next.Value
	prev.next = prev.next.next

	l.size--
	return value, true
}

func (l *SinglyLinkedList[T]) DeleteAfter(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, false
	}

	curr := l.head
	for curr != nil && index > 0 {
		curr = curr.next
		index--
	}

	if curr == nil || curr.next == nil {
		return zero, false
	}

	value := curr.next.Value
	curr.next = curr.next.next

	if curr.next == nil {
		l.tail = curr
	}

	l.size--
	return value, true
}

func (l *SinglyLinkedList[T]) DeleteAtNode(node *Node[T]) (T, bool) {
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

	prev := l.head
	for prev != nil && prev.next != node {
		prev = prev.next
	}

	if prev == nil || prev.next == nil {
		return zero, false
	}

	value := prev.next.Value
	prev.next = prev.next.next

	l.size--
	return value, true
}

func (l *SinglyLinkedList[T]) DeleteAfterNode(node *Node[T]) (T, bool) {
	var zero T

	if node == nil {
		return zero, false
	}

	curr := l.head
	for curr != nil && curr != node {
		curr = curr.next
	}

	if curr == nil || curr.next == nil {
		return zero, false
	}

	value := curr.next.Value
	curr.next = curr.next.next

	if curr.next == nil {
		l.tail = curr
	}

	l.size--
	return value, true
}

func (l *SinglyLinkedList[T]) DeleteValue(value T) bool {
	if l.head == nil {
		return false
	}

	if l.head.Value == value {
		l.DeleteFirst()
		return true
	}

	prev := l.head
	for prev != nil && prev.next != nil {
		if prev.next.Value == value {
			prev.next = prev.next.next

			if prev.next == nil {
				l.tail = prev
			}

			l.size--
			return true
		}

		prev = prev.next
	}
	return false
}

func (l *SinglyLinkedList[T]) Get(index int) (T, bool) {
	var zero T
	if index < 0 || index >= l.size {
		return zero, false
	}

	curr := l.head
	for curr != nil && index > 0 {
		curr = curr.next
		index--
	}

	if curr == nil {
		return zero, false
	}

	return curr.Value, true
}

func (l *SinglyLinkedList[T]) GetNode(index int) *Node[T] {
	if index < 0 || index >= l.size {
		return nil
	}

	curr := l.head
	for curr != nil && index > 0 {
		curr = curr.next
		index--
	}

	return curr
}

func (l *SinglyLinkedList[T]) Copy() *SinglyLinkedList[T] {
	if l.size == 0 {
		return New[T]()
	}

	out := New[T]()
	l.Traverse(func(index int, value T) { out.AddLast(value) })
	return out
}

func (l *SinglyLinkedList[T]) ToSlice() []T {
	out := make([]T, 0, l.size)
	l.Traverse(func(index int, value T) { out = append(out, value) })
	return out
}

func (l *SinglyLinkedList[T]) Search(value T) int {
	i := 0
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.Value == value {
			return i
		}
		i++
	}
	return -1
}

func (l *SinglyLinkedList[T]) Contains(value T) bool {
	return l.Search(value) != -1
}

func (l *SinglyLinkedList[T]) Equal(other *SinglyLinkedList[T]) bool {
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

func (l *SinglyLinkedList[T]) Traverse(fn func(index int, value T)) {
	i := 0
	for curr := l.head; curr != nil; curr = curr.next {
		fn(i, curr.Value)
		i++
	}
}

func (l *SinglyLinkedList[T]) Reverse() *SinglyLinkedList[T] {
	var prev, next *Node[T]
	out := l.Copy()
	curr := out.head
	for curr != nil {
		next = curr.next
		curr.next = prev

		prev = curr
		curr = next
	}
	out.head, out.tail = out.tail, out.head
	return out
}

func (l *SinglyLinkedList[T]) IsSorted() bool {
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

func (l *SinglyLinkedList[T]) Swap(index1, index2 int) bool {
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
	var prev, prev1, prev2, n1, n2 *Node[T]
	for curr := l.head; curr != nil; curr = curr.next {
		if i == index1 {
			prev1, n1 = prev, curr
		}

		if i == index2 {
			prev2, n2 = prev, curr
			break
		}

		i++
		prev = curr
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if prev1 == nil {
		l.head = n2
	} else {
		prev1.next = n2
	}

	if n1.next == n2 {
		n1.next = n2.next
		n2.next = n1

		if n1.next == nil {
			l.tail = n1
		}

		return true
	}

	prev2.next = n1
	n1.next, n2.next = n2.next, n1.next

	if n1.next == nil {
		l.tail = n1
	} else if n2.next == nil {
		l.tail = n2
	}

	return true
}

func (l *SinglyLinkedList[T]) String() string {
	if l.head == nil {
		return "[nil]"
	}

	var sb strings.Builder
	sb.WriteString("[(head) ")

	for curr := l.head; curr != nil; curr = curr.next {
		fmt.Fprintf(&sb, "%v", curr.Value)
		if curr.next != nil {
			sb.WriteString(" -> ")
		} else {
			sb.WriteString(" (tail) -> nil]")
		}
	}
	return sb.String()
}
