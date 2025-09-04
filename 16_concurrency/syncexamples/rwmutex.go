package syncexamples

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type letterCounter struct {
	// RWMutex allows multiple concurrent readers or one exclusive writer
	sync.RWMutex
	counts map[rune]int
}

func newLetterCounter() *letterCounter {
	return &letterCounter{counts: make(map[rune]int)}
}

func (lc *letterCounter) add(r rune) {
	letter := unicode.ToLower(r)
	if letter < 'a' || letter > 'z' {
		return
	}

	// Acquire write lock for exclusive access to the map
	lc.Lock()
	defer lc.Unlock()

	lc.counts[letter]++
}

func (lc *letterCounter) String() string {
	// Acquire read lock to safely read the map concurrently
	lc.RLock()
	defer lc.RUnlock()

	var builder strings.Builder
	for letter, count := range lc.counts {
		fmt.Fprintf(&builder, "%q: %d, ", letter, count)
	}

	s := builder.String()
	if len(s) > 2 {
		s = s[:len(s)-2]
	}
	return s
}

func RWMutexExample() {
	lc := newLetterCounter()

	texts := []string{
		"Hello, World!",
		"Concurrency in Go is powerful.",
		"Mutexes help protect shared data.",
	}

	var wg sync.WaitGroup

	for _, text := range texts {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			for _, letter := range text {
				lc.add(letter)
			}
		}(text)
	}

	wg.Wait()
	fmt.Println(lc)
}
