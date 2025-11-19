package main

import "fmt"

type BankAccount interface {
	GetBalance() int           // Request the bank account to only provide the money left, type could probably be Float in real life
	Deposit(amount int)        // Request the bank account to deposit, and thereby add the funds, type could be Float in real life
	Withdraw(amount int) error // Request to get money, with a condition, i.e. if there is enough funds in the account
}

func main() {
	wf := NewWellsFargo()

	fmt.Println("New Wells Fargo bank account created..")

	wf.Deposit(100)

	currentBalance := wf.GetBalance()

	fmt.Printf("WF balance: %d\n", currentBalance)

	wf.Withdraw(50)

	currentBalance = wf.GetBalance()

	fmt.Printf("WF balance: %d\n", currentBalance)

	bc := NewBitCoinWallet()

	fmt.Println("New Bitcoin account created..")

	bc.Deposit(10000)
}
