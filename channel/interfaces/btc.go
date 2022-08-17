package main

import "errors"

type BTC struct {
	balance int
	fee int
}

func NewBTC() *BTC {
	return &BTC{
		balance:0,
		fee: 100,
	}
}

func (b *BTC) GetBalance() int {
	return b.balance
}

func (b *BTC) Deposit(amount int) {
	b.balance += amount
}


func (b *BTC) Withdraw(amount int) error {

	nebBalance := b.balance - amount - b.fee

	if nebBalance < 0 {
		return errors.New("no funds")
	}

	b.balance = nebBalance
	return nil
}