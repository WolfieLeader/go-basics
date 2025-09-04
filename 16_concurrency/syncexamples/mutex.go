package syncexamples

import (
	"fmt"
	"sync"
)

type bankAccount struct {
	// Mutex protects both `amount` and `personalFee` to ensure safe concurrent access.
	sync.Mutex
	amount      float64
	personalFee float64
}

func newBankAccount() *bankAccount {
	return &bankAccount{}
}

func (a *bankAccount) deposit(amount float64) {
	// Lock the mutex to ensure exclusive access to the account's state
	a.Lock()
	a.amount += amount * (1 - a.personalFee/100)
	// Ensure the mutex is unlocked even if an error or panic occurs
	a.Unlock()
}

func (a *bankAccount) fee(feePercentage float64) {
	a.Lock()
	defer a.Unlock() // Use defer to ensure it is always unlocked

	if feePercentage < 0 || feePercentage > 100 {
		return
	}
	a.personalFee = feePercentage
}

func (a *bankAccount) balance() float64 {
	a.Lock()
	defer a.Unlock()
	return a.amount
}

func MutexExample() {
	account := newBankAccount()
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Go(func() {
			account.fee((float64(i) + 1) * 2.5)
			for range 10 {
				account.deposit(100)
			}
		})
	}

	wg.Wait()
	fmt.Println("Final balance:", account.balance())
}
