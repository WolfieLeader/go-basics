package main

import (
	"fmt"

	arraydeque "github.com/WolfieLeader/go-basics/dsa/data_structures/queue/array_deque"
	arrayqueue "github.com/WolfieLeader/go-basics/dsa/data_structures/queue/array_queue"
	listdeque "github.com/WolfieLeader/go-basics/dsa/data_structures/queue/linked_list_deque"
	listqueue "github.com/WolfieLeader/go-basics/dsa/data_structures/queue/linked_list_queue"
)

func main() {
	fmt.Println("Array Queue Example:")
	arrayQueueExample()
	fmt.Println()

	fmt.Println("Linked List Queue Example:")
	linkedListQueueExample()
	fmt.Println()

	fmt.Println("Array Deque Example:")
	arrayDequeExample()
	fmt.Println()

	fmt.Println("Linked List Deque Example:")
	linkedListDequeExample()
	fmt.Println()
}

func arrayQueueExample() {
	q := arrayqueue.New[int]()
	q.Enqueue(1, 2, 3)
	fmt.Printf("- Queue: %v\n", q)
	q.Dequeue()
	fmt.Printf("- Dequeued: %v\n", q)
	for i := range 10 {
		q.Enqueue(i)
	}
	fmt.Printf("- Enqueued: %v\n", q)
	v, _ := q.Peek()
	fmt.Printf("- Peeked: %v\n", v)
	for range 5 {
		q.Dequeue()
	}
	fmt.Printf("- Dequeued: %v\n", q)
	v, _ = q.Peek()
	fmt.Printf("- Peeked: %v\n", v)
	q.Clear()
	c := q.Copy()
	c.Enqueue(1, 2, 3, 4)
	fmt.Printf("- Are queues equal? %v == %v: %t\n", q, c, q.Equal(c))
	q.Enqueue(1, 2, 3, 4)
	fmt.Printf("- Are queues equal? %v == %v: %t\n", q, c, q.Equal(c))
}

func linkedListQueueExample() {
	q := listqueue.New[int]()
	q.Enqueue(1, 2, 3)
	fmt.Printf("- Queue: %v\n", q)
	q.Dequeue()
	fmt.Printf("- Dequeued: %v\n", q)
	for i := range 10 {
		q.Enqueue(i)
	}
	fmt.Printf("- Enqueued: %v\n", q)
	v, _ := q.Peek()
	fmt.Printf("- Peeked: %v\n", v)
	for range 5 {
		q.Dequeue()
	}
	fmt.Printf("- Dequeued: %v\n", q)
	v, _ = q.Peek()
	fmt.Printf("- Peeked: %v\n", v)
	q.Clear()
	c := q.Copy()
	c.Enqueue(1, 2, 3, 4)
	fmt.Printf("- Are queues equal? %v == %v: %t\n", q, c, q.Equal(c))
	q.Enqueue(1, 2, 3, 4)
	fmt.Printf("- Are queues equal? %v == %v: %t\n", q, c, q.Equal(c))
}

func arrayDequeExample() {
	d := arraydeque.New[int]()
	d.EnqueueLast(1, 2, 3)
	fmt.Printf("- Queue: %v\n", d)
	d.DequeueFirst()
	fmt.Printf("- Dequeued: %v\n", d)
	for i := range 10 {
		d.EnqueueFirst(i)
	}
	fmt.Printf("- Enqueued: %v\n", d)
	v, _ := d.PeekLast()
	fmt.Printf("- Peeked: %v\n", v)
	for range 5 {
		d.DequeueFirst()
	}
	fmt.Printf("- Dequeued: %v\n", d)
	v, _ = d.PeekFirst()
	fmt.Printf("- Peeked: %v\n", v)
	d.Clear()
	c := d.Copy()
	c.EnqueueFirst(1, 2, 3, 4)
	fmt.Printf("- Are deques equal? %v == %v: %t\n", d, c, d.Equal(c))
	d.EnqueueFirst(1, 2, 3, 4)
	fmt.Printf("- Are deques equal? %v == %v: %t\n", d, c, d.Equal(c))
}

func linkedListDequeExample() {
	d := listdeque.New[int]()
	d.EnqueueLast(1, 2, 3)
	fmt.Printf("- Queue: %v\n", d)
	d.DequeueFirst()
	fmt.Printf("- Dequeued: %v\n", d)
	for i := range 10 {
		d.EnqueueFirst(i)
	}
	fmt.Printf("- Enqueued: %v\n", d)
	v, _ := d.PeekLast()
	fmt.Printf("- Peeked: %v\n", v)
	for range 5 {
		d.DequeueFirst()
	}
	fmt.Printf("- Dequeued: %v\n", d)
	v, _ = d.PeekFirst()
	fmt.Printf("- Peeked: %v\n", v)
	d.Clear()
	c := d.Copy()
	c.EnqueueFirst(1, 2, 3, 4)
	fmt.Printf("- Are deques equal? %v == %v: %t\n", d, c, d.Equal(c))
	d.EnqueueFirst(1, 2, 3, 4)
	fmt.Printf("- Are deques equal? %v == %v: %t\n", d, c, d.Equal(c))
}
