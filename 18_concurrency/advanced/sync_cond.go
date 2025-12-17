package advanced

import (
	"fmt"
	"sync"
	"time"
)

type boundedQueue[T any] struct {
	mu       sync.Mutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
	items    []T
	limit    int
	closed   bool
}

func newQueue[T any](cap int) *boundedQueue[T] {
	q := &boundedQueue[T]{
		items: make([]T, 0, cap),
		limit: cap,
		// mu is zero-value initialized
	}
	// Both conditions use the same mutex
	q.notEmpty = sync.NewCond(&q.mu)
	q.notFull = sync.NewCond(&q.mu)
	return q
}

func (q *boundedQueue[T]) Enqueue(item T) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Wait in a loop while the queue is full
	for len(q.items) == q.limit && !q.closed {
		q.notFull.Wait()
	}

	if q.closed {
		return false
	}

	q.items = append(q.items, item)
	q.notEmpty.Signal() // Wake up one waiting Dequeue
	return true
}

func (q *boundedQueue[T]) Dequeue() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Wait in a loop while the queue is empty
	for len(q.items) == 0 && !q.closed {
		q.notEmpty.Wait()
	}

	if len(q.items) == 0 && q.closed {
		var zero T
		return zero, false
	}

	v := q.items[0]
	q.items = q.items[1:]
	q.notFull.Signal() // Wake up one waiting Enqueue
	return v, true
}

func (q *boundedQueue[T]) Close() {
	q.mu.Lock()
	q.closed = true
	q.notEmpty.Broadcast() // Wake up all waiting Dequeue
	q.notFull.Broadcast()  // Wake up all waiting Enqueue
	q.mu.Unlock()
}

func SyncCondExample() {
	var (
		q             = newQueue[int](3)
		prodWg        sync.WaitGroup
		consWg        sync.WaitGroup
		producerCount = 2
		itemsPerProd  = 5
		consumerCount = 3
	)

	// Producers
	prodWg.Add(producerCount)
	for p := 1; p <= producerCount; p++ {
		go func(p int) {
			defer prodWg.Done()
			for i := 1; i <= itemsPerProd; i++ {
				v := p*100 + i
				if !q.Enqueue(v) {
					return
				}
				fmt.Printf("🔴 [Producer %d] Enqueued: %d\n", p, v)
				time.Sleep(75 * time.Millisecond)
			}
		}(p)
	}
	go func() { prodWg.Wait(); q.Close() }()

	// Consumers
	consWg.Add(consumerCount)
	for c := 1; c <= consumerCount; c++ {
		go func(c int) {
			defer consWg.Done()
			for {
				v, ok := q.Dequeue()
				if !ok {
					return
				}
				fmt.Printf("🔵 [Consumer %d] Dequeued: %d\n", c, v)
				time.Sleep(100 * time.Millisecond)
			}
		}(c)
	}
	consWg.Wait()

	fmt.Println("- All producers and consumers completed!")
}
