package main

import "fmt"

func main() {

	var age int

	fmt.Print("Write your age: ")
	_, err := fmt.Scanln(&age)

	if err != nil {
		fmt.Println("Your input have incorrect type")
		return
	}

	switch {

	case age < 13:
		fmt.Println("\nChild")

	case age < 18:
		fmt.Println("\nTeenager")

	case age < 60:
		fmt.Println("\nAdult")

	default:
		fmt.Println("\nSenior")

	}

}
