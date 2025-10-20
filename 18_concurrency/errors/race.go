package errors

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func RaceConditionExample() {
	fmt.Println("\nRace Condition Example:")
	var wg sync.WaitGroup
	const n = 1000

	// ❌ Racy counter (commented out to keep output clean)
	var racy int
	wg.Add(n)
	for range n {
		go func() {
			racy++ // data race: write without synchronization
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("- racy counter:", racy) // This may print a value less than n due to data race
	// Run: `go run -race .` to see detector catch it

	// ✅ Fixed using atomic
	var safe int64
	wg.Add(n)
	for range n {
		go func() {
			atomic.AddInt64(&safe, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("- safe counter:", atomic.LoadInt64(&safe))
}
