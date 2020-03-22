package explore

import "sync"

// Account .
type Account struct {
	name   string
	amount uint32
	mutex  sync.Mutex
}

// Deposit .
func (a *Account) Deposit(amount uint32) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.amount = a.amount + amount
}

// Balance .
func (a *Account) Balance() uint32 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.amount
}
