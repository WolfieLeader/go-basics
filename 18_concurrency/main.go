package main

import (
	basic "github.com/WolfieLeader/go-basics/18_concurrency/basics"
	"github.com/WolfieLeader/go-basics/18_concurrency/errors"
	"github.com/WolfieLeader/go-basics/18_concurrency/patterns"
	"github.com/WolfieLeader/go-basics/18_concurrency/sync"
)

func main() {
	basic.GoroutineExample()
	basic.IgnoredGoroutineExample()
	basic.ChannelExample()
	basic.ChannelIterationExample()
	basic.ChannelDirectionExample()
	basic.SelectExample()
	basic.SelectSendExample()
	basic.ContextTimeoutExample()
	basic.ContextCancellationExample()
	basic.ContextValuesExample()
	basic.IoPipeExample()
	sync.WaitGroupExample()
	sync.ModernWaitGroupExample()
	sync.WaitGroupFetchExample()
	sync.MutexExample()
	sync.RWMutexExample()
	sync.OnceExample()
	sync.PoolExample()
	sync.CondExample()
	sync.AtomicExample()
	errors.DeadlockExample()
	errors.RaceConditionExample()
	errors.GoroutineLeakExample()
	patterns.FanOutFanInExample()
	patterns.GeneratorExample()
	patterns.PipelineExample()
	patterns.WorkerPoolExample()
	patterns.SemaphoreExample()
}
