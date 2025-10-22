package main

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/20_advanced_concurrency/errors"
	"github.com/WolfieLeader/go-basics/20_advanced_concurrency/patterns"
	"github.com/WolfieLeader/go-basics/20_advanced_concurrency/sync"
)

func main() {
	fmt.Println("io.Pipe Example:")
	ioPipeExample()
	fmt.Println()

	fmt.Println("Generator Example:")
	patterns.GeneratorExample()
	fmt.Println()

	fmt.Println("Pipeline Example:")
	patterns.PipelineExample()
	fmt.Println()

	fmt.Println("Fan-Out / Fan-In Example:")
	patterns.FanOutFanInExample()
	fmt.Println()

	fmt.Println("Worker Pool Example:")
	patterns.WorkerPoolExample()
	fmt.Println()

	fmt.Println("Semaphore Example:")
	patterns.SemaphoreExample()
	fmt.Println()

	fmt.Println("Sync Once Example")
	sync.OnceExample()
	fmt.Println()

	fmt.Println("Sync RWMutex Example")
	sync.RWMutexExample()
	fmt.Println()

	fmt.Println("Sync Pool Example")
	sync.PoolExample()
	fmt.Println()

	fmt.Println("Sync Cond Example")
	sync.CondExample()
	fmt.Println()

	fmt.Println("Sync Atomic Example")
	sync.AtomicExample()
	fmt.Println()

	fmt.Println("Race Condition Example")
	errors.RaceConditionExample()
	fmt.Println()

	fmt.Println("Goroutine Leak Example")
	errors.GoroutineLeakExample()
	fmt.Println()

	fmt.Println("Deadlock Example")
	errors.DeadlockExample()
	fmt.Println()
}
