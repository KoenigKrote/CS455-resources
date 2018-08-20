package account

import (
	"fmt"
	"sync"
)

type Account struct {
	balance float64
	sync.Mutex
}

func New() *Account {
	return &Account{
		balance: 0,
	}
}

func (a Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	a.balance -= amount
}

func (a *Account) SafeDeposit(amount float64) {
	a.Lock()
	a.balance += amount
	a.Unlock()
}

func (a *Account) SafeWithdraw(amount float64) {
	a.Lock()
	a.balance -= amount
	a.Unlock()
}

func (a Account) String() string {
	return fmt.Sprintf("Balance: %.2f", a.balance)
}
