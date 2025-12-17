package advanced

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func AtomicExample() {
	var (
		counter        int64
		wg             sync.WaitGroup
		start          = time.Now()
		workerCount    = 5
		tasksPerWorker = 1500
	)

	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()

			for j := 1; j <= tasksPerWorker; j++ {
				// Atomic operations ensure safe concurrent access to the counter
				atomic.AddInt64(&counter, 1)

				if j%500 == 0 {
					fmt.Printf("- [Worker %d]: Processed task %d\n", w, j/500)
				}

				time.Sleep(time.Millisecond)
			}
		}(w)
	}

	wg.Wait()

	final := atomic.LoadInt64(&counter)
	fmt.Printf("✅ All workers done. Total tasks: %d (expected: %d)\n", final, int64(workerCount*tasksPerWorker))
	fmt.Printf("⏱️ Elapsed: %s\n", time.Since(start).Round(time.Millisecond))
}
