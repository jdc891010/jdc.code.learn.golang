package main

import (
	"errors"
	"fmt"
)

type BitcoinWallet struct {
	balance        int
	transactionFee int
}

func NewBitCoinWallet() *BitcoinWallet {
	return &BitcoinWallet{balance: 100, transactionFee: 1} // starting with finders bonus
}

func (bc *BitcoinWallet) GetBalance() int {
	return bc.balance
}

func (bc *BitcoinWallet) Deposit(amount int) {
	fmt.Printf("Deposit of value %d received\n", amount)
}

func (bc *BitcoinWallet) Withdraw(amount int) error {
	if amount > bc.balance {
		return errors.New("Insufficent funds")
	} else {
		bc.balance -= (amount + int(bc.transactionFee))
		fmt.Printf("Amount: %d has been withdrawn, balance is now %d\n", amount, bc.balance)
		return nil
	}
}
