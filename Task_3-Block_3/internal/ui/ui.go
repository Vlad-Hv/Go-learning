package ui

import (
	"fmt"
)

func GetOption()(int, error){
	var option int
	printMenu()
	_, err := fmt.Scanln(&option)
	return option, err
}

func printMenu(){
	fmt.Println("---Menu---")
	fmt.Println("1. Show all events\n2. Show event statistics\n3. Show only errors\n4. Find the most frequent event\n5. Exit\nChoose the option:")
}