package ui

import (
	"fmt"
)

func GetMenuOption() (int, error) {
	var option int
	printMenu()
	_, err := fmt.Scanln(&option)
	return option, err
}

func printMenu() {
	fmt.Println("---Menu---")
	fmt.Println("1. Create membership\n2. Show memberships\n3. Use one visit\n4. Freeze membership\n5. Exit")
	fmt.Print("Choose option:")
}
