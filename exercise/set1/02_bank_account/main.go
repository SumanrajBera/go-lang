package main

import (
	"errors"
	"fmt"
)

type BankService interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
}

type Account struct {
	name    string
	balance float64
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
	fmt.Println("Your current balance:", acc.balance)
}

func (acc *Account) Withdraw(amount float64) error {
	if amount < 0 {
		return errors.New("Amount can't be negative")
	}

	if acc.balance < amount {
		return errors.New("Amount you have entered exceeds the limit")
	}

	acc.balance -= amount
	fmt.Println("Your current balance after withdrawal:", acc.balance)
	return nil
}

func main() {
	myAcc := Account{name: "Sumanraj", balance: 2500}
	myAcc.Deposit(80)
	err := myAcc.Withdraw(2590)
	if err != nil {
		fmt.Println(err)
	}

	err = myAcc.Withdraw(500)
	if err != nil {
		fmt.Println(err)
	}

	err = myAcc.Withdraw(-500)
	if err != nil {
		fmt.Println(err)
	}
}
