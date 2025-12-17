package advanced

import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)

type scratch struct {
	buf []byte
}

// sync.Pool is used to manage a pool of reusable objects to reduce memory allocations.
var scratchPool = sync.Pool{
	New: func() any {
		fmt.Println("✅ Allocating new scratch...")
		return &scratch{buf: make([]byte, 1024)}
	},
}

func SyncPoolExample() {
	var (
		workerCount    = 3
		tasksPerWorker = 3
		wg             sync.WaitGroup
	)

	for w := 1; w <= workerCount; w++ {
		w := w
		wg.Go(func() {
			for t := 1; t <= tasksPerWorker; t++ {
				s := scratchPool.Get().(*scratch)
				s.buf = s.buf[:cap(s.buf)] // Reset slice length

				// Simulate work by filling the buffer with random data
				rand.Read(s.buf)
				fmt.Printf("- [Worker %d, Task %d]: used %d bytes\n", w, t, len(s.buf))

				// Return to pool for reuse
				scratchPool.Put(s)
				fmt.Printf("< [Worker %d, Task %d]: returned scratch to pool\n\n", w, t)
				time.Sleep(80 * time.Millisecond)
			}
		})
		time.Sleep(10 * time.Millisecond) // Allocate less buffers concurrently
	}

	wg.Wait()
	fmt.Println("All workers done.")
}
