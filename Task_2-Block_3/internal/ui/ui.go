package ui

import (
	"fmt"
)

func GetExpenceInfo() (string, int, error) {
	var category string
	var amount int
	fmt.Println("Enter category and amount:")
	_, err := fmt.Scanln(&category, &amount)
	return category, amount, err
}

func GetMenuOption() (int, error) {
	var option int
	printMenu()
	_, err := fmt.Scanln(&option)
	return option, err
}

func printMenu() {
	fmt.Println("---Menu---")
	fmt.Println("1. Add expense\n2. Show expenses\n3. Show total amount\n4. Exit")
	fmt.Println("Choose the option:")
}
