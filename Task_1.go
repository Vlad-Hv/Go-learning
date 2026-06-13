package main

import "fmt"

func main() {
	const username = "admin"
	const password = "VAD216"

	var userName string
	var userPassword string

	fmt.Println("Write your username and password to log in")

	fmt.Scanln(&userName)
	fmt.Scanln(&userPassword)

	if userName == username && userPassword == password {
		fmt.Println("\nLogin successful")
	} else if userName == username && userPassword != password {
		fmt.Println("\nincorect password")
	} else if userName != username && userPassword == password {
		fmt.Println("\nIncorrect Name")
	} else {
		fmt.Println("Invalid credentials")
	}

}
