package main

import (
	"errors"
	"fmt"
)

func main() {
	var balance float64 = 1000

	for {
		fmt.Println("\n===== Mini Bank =====\n\n ")
		option, err := menuOption()

		if err != nil {
			fmt.Println(err)
			return
		}

		if option == 4 {
			break
		}

		switch option {

		case 1:
			showBalance(balance)

		case 2:
			deposit, err := depositMoney(balance)

			if err != nil {
				fmt.Println(err)
				return
			}

			balance = deposit
			success()

		case 3:
			result, err := withdrawMoney(balance)

			if err != nil {
				fmt.Println(err)
				return
			}

			balance = result
			success()

		default:
			fmt.Println("\nWrong option, try again")
			continue

		}
	}

	fmt.Println("\n\nGoodbye :)")

}

func getUserAnswer() (int, error) {
	var userAnswer int
	fmt.Print("Your budjet = 1000$\nDo you want to start? (1 - y/n - 2): ")
	_, err := fmt.Scanln(&userAnswer)

	if err != nil {
		return 0, errors.New("Incorrect input type")
	}

	return userAnswer, nil
}

func menuOption() (int, error) {
	var option int
	fmt.Print("\n1. Show balance\n2. Deposit money\n3. Withdraw money\n4. Exit\n\nChoose option: ")
	_, err := fmt.Scanln(&option)

	if err != nil {
		return 0, errors.New("entered option must have integer type")
	}

	return option, nil
}

func success() {
	fmt.Println("Balance changed successfully!")
}

func showBalance(balance float64) {
	fmt.Println("Current balance:", balance)
}

func depositMoney(balance float64) (float64, error) {
	var depSum float64
	fmt.Print("Enter sum which you wanna deposit: ")
	fmt.Scanln(&depSum)

	if depSum <= 0 {
		return 0, fmt.Errorf("incorrect deposit sum: %v", depSum)
	}

	return depSum + balance, nil
}

func withdrawMoney(balance float64) (float64, error) {
	var withdraw float64
	fmt.Print("\n\nEnte Withdraw sum: ")
	fmt.Scanln(&withdraw)

	if withdraw <= 0 {
		return 0, fmt.Errorf("withdraw sum must be positive: %v", withdraw)
	}

	if withdraw > balance {
		return 0, errors.New("not enough money")
	}

	return balance - withdraw, nil
}
