package advanced

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type config struct {
	value int
}

var (
	cfg  *config
	once sync.Once
)

func loadConfig(goroutineId int) {
	once.Do(func() {
		fmt.Printf("⚙️ [Goroutine %d]: Loading configuration...\n", goroutineId)
		time.Sleep(100 * time.Millisecond) // Simulate loading time
		cfg = &config{value: goroutineId * 100}
		fmt.Println("✅ Configuration loaded.")
	})
}

func SyncOnceExample() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(200)) * time.Microsecond)
			loadConfig(i)
			fmt.Printf("- Goroutine %d: Got this config value: %d\n", i, cfg.value)
		}(i)
	}

	wg.Wait()
	fmt.Println("- All goroutines completed.")
}
