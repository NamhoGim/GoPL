package bank

import (
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	// Deposit [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			Withdraw(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	if got, balance := Withdraw(1), Balance(); got || balance != 0 {
		t.Errorf("Balance = %d, Withdraw result: %v", balance, got)
	}
}
