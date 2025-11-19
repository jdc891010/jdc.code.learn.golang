package main

import (
	"errors"
	"fmt"
)

type WellsFargo struct {
	balance        int
	transactionFee int
} // technically the only thing that the bank account needs, the rest are transactions on it

func NewWellsFargo() *WellsFargo {
	return &WellsFargo{balance: 0, transactionFee: 2} // returns a memory reference to the wf struct defined above, with init value
} // Returns a pointer to the account, which could updated

func (w *WellsFargo) GetBalance() int {
	return w.balance
}

func (w *WellsFargo) Deposit(amount int) {
	fmt.Printf("Deposit of value %d received\n", amount)

	w.balance += amount // Using the pointer to update the account
}

func (w *WellsFargo) Withdraw(amount int) error {
	if amount > w.balance {
		return errors.New("Insufficient funds") // i.e. request invalid, not enough money
	} else {
		w.balance -= amount
		fmt.Printf("Amount: %d has been withdrawn, balance is now: %d\n", amount, w.balance)
		return nil
	}
}

// Remember that the *WellsFargo, is still a type, and it needs to be of type *WellsFargo
// hence the return from NewWellsFargo is reference to the memory address of newly crated struct
