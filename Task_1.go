package main

import (
	"errors"
	"fmt"
)

func main() {
	var balance float64 = 1000.00

	err := changeBalance(&balance)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Balance:", balance)
}

func changeBalance(balance *float64) error {
	var deposit float64

	fmt.Println("Enter money to deposit: ")
	_, err := fmt.Scanln(&deposit)

	if err != nil {
		return errors.New("invalid type")
	}

	if deposit <= 0 {
		return errors.New("deposit mustnot be less than 1")
	}

	*balance += deposit

	return nil
}
