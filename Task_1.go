package main

import "fmt"

func main() {
	username, userAge, err := printUserInfo()

	if err != nil {
		return
	}

	isAdult := isAdult(userAge)

	printInfo(username, userAge, isAdult)
}

func printUserInfo() (string, int, error) {
	var username string
	var userAge int

	fmt.Println("Write your name:")
	fmt.Scanln(&username)
	fmt.Println("Well, write your age:")
	_, err := fmt.Scanln(&userAge)

	return username, userAge, err
}

func isAdult(userAge int) bool {
	return userAge >= 18
}

func printInfo(username string, userAge int, isAdult bool) {
	fmt.Println("user", username, "is", userAge, "years old")
	if !isAdult {
		fmt.Println("Access denied")
	} else {
		fmt.Println("Access granted")
	}
}
