package main

import "fmt"

type IBankAccount interface {
	GetBalance() int           // Request the bank account to only provide the money left, type could probably be Float in real life
	Deposit(amount int)        // Request the bank account to deposit, and thereby add the funds, type could be Float in real life
	Withdraw(amount int) error // Request to get money, with a condition, i.e. if there is enough funds in the account
} // Using "I" prefix as good way of indicating it is a interface

func main() {
	// Code Version 1:
	// wf := NewWellsFargo()
	// fmt.Println("New Wells Fargo bank account created..")
	// wf.Deposit(100)
	// currentBalance := wf.GetBalance()
	// fmt.Printf("WF balance: %d\n", currentBalance)
	// wf.Withdraw(50)
	// currentBalance = wf.GetBalance()
	// fmt.Printf("WF balance: %d\n", currentBalance)
	// bc := NewBitCoinWallet()
	// fmt.Println("New Bitcoin account created..")
	// bc.Deposit(10000)

	// Code Version 2:
	// Creating a list of bank accounts, treating the bitcoin and wf as the same thing, comply with inteface blueprint
	myAccounts := []IBankAccount{
		NewBitCoinWallet(),
		NewWellsFargo(),
	}

	for _, account := range myAccounts {
		account.Deposit(100)

		err := account.Withdraw(70)
		if err != nil { // Reminding myself of the go err pattern, checking if it is not nil, i.e. something did produce an error
			fmt.Printf("Erorr occured: %s", err)
		} else {
			fmt.Print("Withdrawl success..\n")
		}

		balance := account.GetBalance() // Since both are interface type, and they align with the declared types, they have these methods

		fmt.Printf("balance = %d\n", balance) // Using %d, since it is a number
	}
}

// Running the code:
// Std Output
// Deposit of value 100 received
// Amount: 70 has been withdrawn, balance is now 29
// Withdrawl success..
// balance = 29
// Deposit of value 100 received
// Amount: 70 has been withdrawn, balance is now: 30
// Withdrawl success..
// balance = 30
