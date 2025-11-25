package main

import (
	"fmt"

	"github.com/WolfieLeader/go-basics/17_concurrency/advanced"
	"github.com/WolfieLeader/go-basics/17_concurrency/basic"
	"github.com/WolfieLeader/go-basics/17_concurrency/patterns"
)

func main() {
	fmt.Println("Goroutine Example:")
	basic.GoroutineExample()
	fmt.Println()

	fmt.Println("Unbuffered Channel Example:")
	basic.UnbufferedChannelExample()
	fmt.Println()

	fmt.Println("Buffered Channel Example:")
	basic.BufferedChannelExample()
	fmt.Println()

	fmt.Println("Comma-Ok Channel Example:")
	basic.CommaOkChannelExample()
	fmt.Println()

	fmt.Println("Channel Directions Example:")
	basic.ChannelDirectionsExample()
	fmt.Println()

	fmt.Println("Select Example:")
	basic.SelectExample()
	fmt.Println()

	fmt.Println("Send Select Example:")
	basic.SelectSendExample()
	fmt.Println()

	fmt.Println("Sync Wait Group Example:")
	basic.SyncWaitGroupExample()
	fmt.Println()

	fmt.Println("Modern Sync Wait Group Example:")
	basic.ModernSyncWaitGroupExample()
	fmt.Println()

	fmt.Println("Sync Mutex Example:")
	basic.SyncMutexExample()
	fmt.Println()

	fmt.Println("Patterns: Generator")
	patterns.GeneratorExample()
	fmt.Println()

	fmt.Println("Patterns: Pipeline")
	patterns.PipelineExample()
	fmt.Println()

	fmt.Println("Patterns: Fan-Out Fan-In")
	patterns.FanOutFanInExample()
	fmt.Println()

	fmt.Println("Advanced: io.Pipe")
	advanced.IoPipeExample()
	fmt.Println()
}
