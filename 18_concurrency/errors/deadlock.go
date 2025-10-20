package errors

import (
	"fmt"
	"sync"
	"time"
)

func DeadlockExample() {
	fmt.Println("\nDeadlock Example (shown safely):")

	var muA, muB sync.Mutex
	done := make(chan struct{})

	// ❌
	// This would deadlock if both goroutines lock in opposite order.
	// We avoid the actual deadlock by timing out.
	go func() {
		defer close(done)
		muA.Lock()
		defer muA.Unlock()
		time.Sleep(50 * time.Millisecond)
		muB.Lock()
		muB.Unlock()
	}()

	go func() {
		muB.Lock()
		defer muB.Unlock()
		time.Sleep(50 * time.Millisecond)
		muA.Lock()
		muA.Unlock()
	}()

	select {
	case <-done:
		fmt.Println("- finished (no deadlock occurred this run)")
	case <-time.After(150 * time.Millisecond):
		fmt.Println("- potential deadlock detected: opposite lock ordering")
	}

	// ✅ Fix: always acquire locks in the same order (A then B)
	func() {
		muA.Lock()
		defer muA.Unlock()
		muB.Lock()
		defer muB.Unlock()
	}()
	fmt.Println("- fixed by consistent lock ordering")
}
