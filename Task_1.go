package main

import "fmt"

func main() {
	const password = "VAD216"

	var userPassword string

	fmt.Println("Write your password:")
	_, err := fmt.Scanln(&userPassword)

	if err != nil {
		fmt.Println("Passsword must include laters and numbers")
		return
	} else if password == userPassword {
		fmt.Println("\nAccess granted")
	} else {
		fmt.Println("\nAccess denied")
	}

}
