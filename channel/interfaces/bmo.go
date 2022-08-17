package main

import "errors"

type BMO struct {
	balance int
}

func NewBMO() *BMO {
	return &BMO{
		balance:0,
	}
}

func (w *BMO) GetBalance() int {
	return w.balance
}

func (w *BMO) Deposit(amount int) {
	w.balance += amount
}


func (w *BMO) Withdraw(amount int) error {

	newBalance := w.balance - amount

	if newBalance < 0 {
		return errors.New("no funds")
	}

	w.balance = newBalance
	return nil
}