package ui

import (
	"fmt"
)

func GetName() string {
	var username string
	printName()
	fmt.Scanln(&username)
	return username
}

func GetLevel() int {
	var userLevel int
	printLevel()
	fmt.Scanln(&userLevel)
	return userLevel
}

func printName() {
	fmt.Println("Enter your name:")
}

func printLevel() {
	fmt.Println("Enter your level:")
}
