package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type BankAccount struct {
	// Mutex protects both `balance` and `fee` to ensure safe concurrent access.
	sync.Mutex
	balance float64
	fee     float64
}

func NewBankAccount() *BankAccount {
	return &BankAccount{}
}

func (a *BankAccount) Deposit(amount float64) {
	// Lock the mutex to ensure exclusive access to the account's state
	a.Lock()
	// Ensure the mutex is unlocked even if an error or panic occurs
	defer a.Unlock()
	a.balance += amount * (1 - a.fee/100)
}

func (a *BankAccount) Fee(feePercentage float64) {
	a.Lock()
	defer a.Unlock()

	if feePercentage < 0 || feePercentage > 100 {
		return
	}
	a.fee = feePercentage
}

func (a *BankAccount) Balance() float64 {
	a.Lock()
	defer a.Unlock()
	return a.balance
}

func mutexExample() {
	account := NewBankAccount()
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Go(func() {
			account.Fee((float64(i) + 1) * 2.5)
			for range 10 {
				account.Deposit(100)
			}
		})
	}

	wg.Wait()
	fmt.Println("Final balance:", account.Balance())
}

type LetterCounter struct {
	// RWMutex allows multiple concurrent readers or one exclusive writer
	sync.RWMutex
	Counts map[rune]int
}

func NewLetterCounter() *LetterCounter {
	return &LetterCounter{Counts: make(map[rune]int)}
}

func (lc *LetterCounter) Add(r rune) {
	letter := unicode.ToLower(r)
	if letter < 'a' || letter > 'z' {
		return
	}

	// Acquire write lock for exclusive access to the map
	lc.Lock()
	defer lc.Unlock()

	lc.Counts[letter]++
}

func (lc *LetterCounter) String() string {
	// Acquire read lock to safely read the map concurrently
	lc.RLock()
	defer lc.RUnlock()

	var builder strings.Builder
	for letter, count := range lc.Counts {
		fmt.Fprintf(&builder, "%q: %d, ", letter, count)
	}

	s := builder.String()
	if len(s) > 2 {
		s = s[:len(s)-2]
	}
	return s
}

func rwMutexExample() {
	lc := NewLetterCounter()

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
				lc.Add(letter)
			}
		}(text)
	}

	wg.Wait()
	fmt.Println(lc)
}
