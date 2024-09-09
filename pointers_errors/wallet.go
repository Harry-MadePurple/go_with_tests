package pointers_errors

import (
	"errors"
	"fmt"
)

// types
type Bitcoin int

type Stringer interface {
	String() string
}
type Wallet struct {
	balance Bitcoin
}

// Methods
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Functions
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
