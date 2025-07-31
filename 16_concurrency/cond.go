package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxCapacity              = 3
	itemsPerProducer         = 5
	producerAndConsumerCount = 3
)

type Queue struct {
	Data  []string
	Mutex *sync.Mutex
	Cond  *sync.Cond
}

func NewQueue() *Queue {
	var mutex sync.Mutex
	return &Queue{
		// Sets cap to avoid reallocations
		Data:  make([]string, 0, maxCapacity),
		Mutex: &mutex,
		// `sync.Cond` is used to wait for and signal conditions between goroutines.
		// It requires a mutex to protect shared state.
		Cond: sync.NewCond(&mutex),
	}
}

func (q *Queue) Enqueue(item string) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	// Wait until there is space in the queue
	// Uses a loop to handle spurious wakeups
	// This is necessary to ensure that the condition is still valid after waking up
	for len(q.Data) >= maxCapacity {
		q.Cond.Wait()
	}

	q.Data = append(q.Data, item)

	// Signal a waiting consumer that an item is available
	q.Cond.Signal()
}

func (q *Queue) Dequeue() string {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	for len(q.Data) == 0 {
		q.Cond.Wait()
	}

	item := q.Data[0]
	q.Data = q.Data[1:]

	// Signal a waiting producer that space is available
	q.Cond.Signal()

	return item
}

func producer(start time.Time, id int, queue *Queue, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range itemsPerProducer {
		item := fmt.Sprintf("Item-%d%c", i+1, 'A'+byte(id))
		queue.Enqueue(item)
		fmt.Printf("🔴 [%s] P-%c: Produced: %q\n", time.Since(start).Round(time.Millisecond), 'A'+byte(id), item)
		time.Sleep(150 * time.Millisecond)
	}
}

func consumer(start time.Time, id int, queue *Queue) {
	for {
		item := queue.Dequeue()
		fmt.Printf("🔵 [%s] C-%d: Consumed: %q\n", time.Since(start).Round(time.Millisecond), id, item)
		time.Sleep(300 * time.Millisecond)
	}
}

func condExample() {
	var wg sync.WaitGroup
	start := time.Now()
	queue := NewQueue()

	for i := range producerAndConsumerCount {
		// Add a wait group counter for each producer (since consumer depends on producer)
		wg.Add(1)
		go producer(start, i, queue, &wg)
		go consumer(start, i, queue)
	}

	// Waits for producers only
	wg.Wait()
	fmt.Println("Finished all producers and consumers")
}
