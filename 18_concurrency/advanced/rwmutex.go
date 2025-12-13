package advanced

import (
	"fmt"
	"sync"
	"time"
)

type cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *cache) Get(key string) (string, bool) {
	c.mu.RLock()                      // Acquire read lock
	defer c.mu.RUnlock()              // Release read lock
	time.Sleep(15 * time.Millisecond) // Simulate work
	v, ok := c.data[key]
	return v, ok
}

func (c *cache) Set(key, value string) {
	c.mu.Lock()                       // Acquire write lock
	defer c.mu.Unlock()               // Release write lock
	time.Sleep(30 * time.Millisecond) // Simulate work
	c.data[key] = value
}

func RWMutexExample() {
	var (
		cache       = &cache{data: make(map[string]string)}
		writerCount = 3
		readerCount = 5
		done        = make(chan struct{})
		writeWg     sync.WaitGroup
		readWg      sync.WaitGroup
	)

	// Writers
	writeWg.Add(writerCount)
	for w := 1; w <= writerCount; w++ {
		go func(w int) {
			defer writeWg.Done()
			for i := range 3 {
				v := fmt.Sprintf("%c%d", 'A'+byte(i), w)
				cache.Set("x", v)
				fmt.Printf("🔴 [Writer %d] Set: %q\n", w, v)
				time.Sleep(75 * time.Millisecond)
			}
		}(w)
	}
	go func() { writeWg.Wait(); close(done) }()

	// Readers
	readWg.Add(readerCount)
	for r := 1; r <= readerCount; r++ {
		go func(r int) {
			defer readWg.Done()

			ticker := time.NewTicker(50 * time.Millisecond)
			defer ticker.Stop()

			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					if v, ok := cache.Get("x"); ok {
						fmt.Printf("🔵 [Reader %d] Get: %q\n", r, v)
					}
				}
			}
		}(r)
	}

	readWg.Wait()
}
