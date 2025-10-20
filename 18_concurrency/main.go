package main

import (
	errorexamples "github.com/WolfieLeader/go-basics/18_concurrency/errors"
	"github.com/WolfieLeader/go-basics/18_concurrency/patterns"
	syncexamples "github.com/WolfieLeader/go-basics/18_concurrency/sync"
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
	syncexamples.WaitGroupExample()
	syncexamples.ModernWaitGroupExample()
	syncexamples.WaitGroupFetchExample()
	syncexamples.MutexExample()
	syncexamples.RWMutexExample()
	syncexamples.OnceExample()
	syncexamples.PoolExample()
	syncexamples.CondExample()
	syncexamples.AtomicExample()
	contextTimeoutExample()
	contextCancellationExample()
	contextValuesExample()
	errorexamples.DeadlockExample()
	errorexamples.RaceConditionExample()
	errorexamples.GoroutineLeakExample()
	patterns.FanOutFanInExample()
	patterns.GeneratorExample()
	patterns.PipelineExample()
	patterns.WorkerPoolExample()
	patterns.SemaphoreExample()
}
