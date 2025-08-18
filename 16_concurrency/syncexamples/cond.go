package syncexamples

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

type queue struct {
	data  []string
	mutex *sync.Mutex
	cond  *sync.Cond
}

func newQueue() *queue {
	var mutex sync.Mutex
	return &queue{
		// Sets cap to avoid reallocations
		data:  make([]string, 0, maxCapacity),
		mutex: &mutex,
		// `sync.Cond` is used to wait for and signal conditions between goroutines.
		// It requires a mutex to protect shared state.
		cond: sync.NewCond(&mutex),
	}
}

func (q *queue) enqueue(item string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// Wait until there is space in the queue
	// Uses a loop to handle spurious wakeups
	// This is necessary to ensure that the condition is still valid after waking up
	for len(q.data) >= maxCapacity {
		q.cond.Wait()
	}

	q.data = append(q.data, item)

	// Signal a waiting consumer that an item is available
	q.cond.Signal()
}

func (q *queue) dequeue() string {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	for len(q.data) == 0 {
		q.cond.Wait()
	}

	item := q.data[0]
	q.data = q.data[1:]

	// Signal a waiting producer that space is available
	q.cond.Signal()

	return item
}

func producer(start time.Time, id int, queue *queue, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range itemsPerProducer {
		item := fmt.Sprintf("Item-%d%c", i+1, 'A'+byte(id))
		queue.enqueue(item)
		fmt.Printf("🔴 [%s] P-%c: Produced: %q\n", time.Since(start).Round(time.Millisecond), 'A'+byte(id), item)
		time.Sleep(150 * time.Millisecond)
	}
}

func consumer(start time.Time, id int, queue *queue) {
	for {
		item := queue.dequeue()
		fmt.Printf("🔵 [%s] C-%d: Consumed: %q\n", time.Since(start).Round(time.Millisecond), id, item)
		time.Sleep(300 * time.Millisecond)
	}
}

func CondExample() {
	var wg sync.WaitGroup
	start := time.Now()
	queue := newQueue()

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
