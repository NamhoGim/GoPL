package bank

import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

// This function requires that the lock be held.
func deposit(amount int) { balance = balance + amount }

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}
