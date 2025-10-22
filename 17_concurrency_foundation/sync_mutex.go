package main

import (
	"fmt"
	"sync"
)

type Account struct {
	// Mutex to protect access to the account's fields
	mu      sync.Mutex
	balance float64
	fee     float64
}

func (a *Account) Deposit(amount float64) {
	a.mu.Lock() // Lock the mutex to ensure exclusive access
	a.balance += amount * (1 - a.fee/100)
	a.mu.Unlock() // Unlock the mutex
}

func (a *Account) SetFee(fee float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if fee >= 0 && fee <= 100 {
		a.fee = fee
	}
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func syncMutexExample() {
	var wg sync.WaitGroup
	account := &Account{}

	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			account.SetFee(float64(i+1) * 2.5)
			for range 10 {
				account.Deposit(100)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("- Final balance: %.2f\n", account.GetBalance())
}
