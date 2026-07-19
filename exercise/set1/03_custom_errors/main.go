package main

import (
	"errors"
	"fmt"
)

type InsufficientFundsError struct {
	requested float64
	available float64
}

type BankService interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
}

type Account struct {
	name    string
	balance float64
}

func (e InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: requested %.2f, available %.2f", e.requested, e.available)
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
		return InsufficientFundsError{
			requested: amount,
			available: acc.balance,
		}
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
