package main

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	poolWorkers        = 3
	tasksPerPoolWorker = 3
)

type User struct {
	Id   int
	Data []byte
}

func (u *User) String() string {
	return fmt.Sprintf("User %d", u.Id)
}

func poolExample() {
	var wg sync.WaitGroup

	// sync.Pool is used to help reduce GC pressure and repeated memory allocations.
	// The `New` function is called to create a new instance when the pool is empty and a new item is needed.
	var userPool = sync.Pool{
		New: func() any {
			fmt.Println("✅ Allocating new User")
			return &User{Data: make([]byte, 1024)} // 1KB buffer
		},
	}

	for i := range poolWorkers {
		workerId := i + 1

		wg.Go(func() {
			for j := range tasksPerPoolWorker {
				label := fmt.Sprintf("Worker %d%c", workerId, 'A'+byte(j))

				// Get a User from the pool. If the pool is empty, it will call the New function.
				user := userPool.Get().(*User)
				cryptoRand.Read(user.Data)
				fmt.Printf(" - %s using %s\n", label, user)

				time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)

				user.Id = workerId
				user.Data = user.Data[:cap(user.Data)]

				// Return the User back to the pool for reuse.
				// This helps reduce memory allocations and GC pressure.
				userPool.Put(user)
				fmt.Printf("   %s returned User %d to pool\n", label, user.Id)
			}
			time.Sleep(10 * time.Millisecond)
		})
	}

	wg.Wait()
	fmt.Println("All workers done.")
}
