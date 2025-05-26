package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	_balance Bitcoin
}

func (wallet *Wallet) deposit(amount Bitcoin) {
	fmt.Printf("address of balance in deposit is %p \n", &wallet._balance)
	wallet._balance += amount
}

func (wallet *Wallet) balance() Bitcoin {
	return wallet._balance
}

func (wallet *Wallet) withdraw(amount Bitcoin) error {
	if amount > wallet._balance {
		return errors.New("insufficient funds")
	}
	wallet._balance -= amount
	return nil
}

func (bitcoin Bitcoin) string() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}