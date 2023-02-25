package bank

import (
	"errors"
)

type BankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

type Zenith struct {
	Balance int
}

type Fidelity struct {
	Balance int
}

func NewZenith() *Zenith {
	return &Zenith{
		Balance: 0,
	}
}

func (z *Zenith) GetBalance() int {
	return z.Balance
}

func (z *Zenith) Deposit(amount int) {
	z.Balance += amount
}

func (z *Zenith) Withdraw(amount int) error {
	newBalance := z.Balance - amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	z.Balance = newBalance
	return nil
}

func NewFidelity() *Fidelity {
	return &Fidelity{
		Balance: 0,
	}
}

func (f *Fidelity) GetBalance() int {
	return f.Balance
}

func (f *Fidelity) Deposit(amount int) {
	f.Balance += amount
}

func (f *Fidelity) Withdraw(amount int) error {
	newBalance := f.Balance - amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	f.Balance = newBalance
	return nil
}
