package main

import (
	"github.com/WolfieLeader/go-basics/18_concurrency/errors"
	"github.com/WolfieLeader/go-basics/18_concurrency/patterns"
	"github.com/WolfieLeader/go-basics/18_concurrency/sync"
)

func main() {
	goroutineExample()
	ignoredGoroutineExample()
	channelExample()
	channelIterationExample()
	channelDirectionExample()
	selectExample()
	selectSendExample()
	ioPipeExample()
	sync.WaitGroupExample()
	sync.ModernWaitGroupExample()
	sync.WaitGroupFetchExample()
	sync.MutexExample()
	sync.RWMutexExample()
	sync.OnceExample()
	sync.PoolExample()
	sync.CondExample()
	sync.AtomicExample()
	contextTimeoutExample()
	contextCancellationExample()
	contextValuesExample()
	errors.DeadlockExample()
	errors.RaceConditionExample()
	errors.GoroutineLeakExample()
	patterns.FanOutFanInExample()
	patterns.GeneratorExample()
	patterns.PipelineExample()
	patterns.WorkerPoolExample()
	patterns.SemaphoreExample()
}
