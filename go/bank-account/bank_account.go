package account

import "sync"

// Account represents the customer's account
type Account struct {
	isClosed bool
	balance  int64
	mux      sync.RWMutex
}

// Open creates an account with initial balance
func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}
	return &Account{balance: deposit}
}

// Close closes the account
func (a *Account) Close() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	balance, isClosed := a.balance, a.isClosed

	a.balance = 0
	a.isClosed = true

	return balance, !isClosed
}

// Balance returns the current balance
func (a *Account) Balance() (int64, bool) {
	a.mux.RLock()
	defer a.mux.RUnlock()

	return a.balance, !a.isClosed
}

// Deposit handles a negative amount as a withdrawal
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.isClosed || a.balance < -amount {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
