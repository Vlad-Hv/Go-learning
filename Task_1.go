package main

import "fmt"

func main() {

	const adminName = "admin"
	const adminPassword = "VAD216"

	var username string
	var userPassword string
	var userAge int

	fmt.Println("Write your username:")
	fmt.Scanln(&username)

	if username == adminName {
		fmt.Println("Write your admin's password: ")
		_, err := fmt.Scanln(&userPassword)

		if err != nil {
			fmt.Println("\nIncorrect admin's input\n\nAccess Blocked; Reason: Attack Try")
			return
		}

		if userPassword == adminPassword {
			fmt.Println("\n\nAdmin mode enabled\n ") //Last \n just for terminal appearence(just for me)
			return
		} else {
			fmt.Println("\nIncorrect admin's input\n\nAccess Blocked; Reason: Attack Try")
			return
		}
	}

	fmt.Println("\nSucces! We added your name, write your password:")
	fmt.Scanln(&userPassword)

	fmt.Println("\nSucces! We added your password, write your age:")
	_, err := fmt.Scanln(&userAge)

	if err != nil {
		fmt.Println("\nIncorrect age input")
		return
	}

	if userAge < 13 {
		fmt.Println("\nToo young to play")
	} else {
		fmt.Println("\nWelcome to the game!\n\nusername:", username, "\nPassword: ******\nAge:", userAge)
	}
}
