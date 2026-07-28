package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner           string
	Balance         float64
	IsBlocked       bool
	OperationsCount int
}

func main() {
	account := BankAccount{
		Owner:           "CEO",
		Balance:         1683.32,
		IsBlocked:       false,
		OperationsCount: 291,
	}
	amount, err := getAmount()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = account.Deposit(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	account.PrintStatus()

	err = account.Withdraw(amount)

	if err != nil {
		fmt.Println(err)
		return
	}

	account.PrintStatus()

}

func (account *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	if account.Block() != false {
		return errors.New("account blocked")
	}

	account.Balance += amount
	account.OperationsCount++
	return nil
}

func (account BankAccount) CanWithdraw(amount float64) error {
	if amount < 0 {
		return errors.New("invalid money amount")
	}
	if account.Balance < amount {
		return errors.New("not enough money on the balance")
	}

	if account.Block() != false {
		return errors.New("account blocked")
	}
	return nil
}

func (account BankAccount) Block() bool {
	return account.IsBlocked
}

func (account *BankAccount) Withdraw(amount float64) error {
	err := account.CanWithdraw(amount)

	if err != nil {
		return err
	}

	account.Balance -= amount
	return nil
}

func (account *BankAccount) PrintStatus() error {
	fmt.Println(*account)
	return nil
}

func getAmount() (float64, error) {
	var amount float64
	fmt.Println("Enter amount of money:")
	_, err := fmt.Scanln(&amount)

	if err != nil {
		return 0, errors.New("invalid type")
	}
	return amount, nil
}
